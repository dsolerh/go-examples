package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// request()
	doc, err := request()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	f, err := os.Create("cuballama.current.html")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	err = html.Render(f, doc)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	resultsF, err := os.Create("cuballama.results.html")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	NewHTMLElement(doc).Traverse(func(element *HTMLElement) TraverseGuide {
		if element.Is(atom.Div) && element.HasClasses("Card") {
			fmt.Printf("element: %s\n", element)
			// fmt.Printf("element parent: %s\n", element.Parent())

			content := element.FirstChild().Find(func(el *HTMLElement) bool {
				return el.Is(atom.H4) && el.Parent().Is(atom.Div) && el.Parent().HasClasses("text-box")
			})

			fmt.Printf("h4 content: %s\n", content.Content())

			err := html.Render(resultsF, element.n)
			if err != nil {
				fmt.Println(err)
			}

			return TraverseGuide{NextSibling: true}
		}
		return TraverseGuide{NextSibling: true, FirstChild: true}
	})
}

func request() (*html.Node, error) {
	data := url.Values{}
	data.Set("attr_start_date", "11/02/2025")
	data.Set("attr_end_date", "17/02/2025")
	data.Set("attr_destination_iata", "VRO")
	data.Set("attr_destination_name", "Matanzas")

	reqCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, "POST", "https://www.cuballama.com/viajes/activity/search", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	// add the headers
	setHeaders(
		req,
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Language: en-US,en;q=0.9",
		"Cache-Control: no-cache",
		"Connection: keep-alive",
		"Content-Type: application/x-www-form-urlencoded",
		// "Cookie: PHPSESSID=a4cfdf069f228ced38ebf977bc39952e; wmsession=d0d1c396-284a-439d-915a-b2dd47c864ce-1732145711047; wm_lang_code=en",
		// "DNT: 1",
		"Origin: https://www.cuballama.com",
		"Pragma: no-cache",
		"Referer: https://www.cuballama.com/viajes/activity/search",
		"Sec-Fetch-Dest: document",
		"Sec-Fetch-Mode: navigate",
		"Sec-Fetch-Site: same-origin",
		"Upgrade-Insecure-Requests: 1",
		"User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
		`sec-ch-ua: "Not;A=Brand";v="24", "Chromium";v="128"`,
		"sec-ch-ua-mobile: ?0",
		`sec-ch-ua-platform: "macOS"`,
	)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status %s", resp.Status)
	}

	// b, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("body: %s\n", b)

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing the response body: %w", err)
	}

	return doc, nil
	// return nil, nil
}

func setHeaders(rq *http.Request, headers ...string) {
	for _, headerPair := range headers {
		keyvalue := strings.Split(headerPair, ": ")
		rq.Header.Set(keyvalue[0], keyvalue[1])
	}
}

type TraverseGuide struct {
	NextSibling bool
	FirstChild  bool
}

type HTMLElement struct {
	n *html.Node
}

func NewHTMLElement(node *html.Node) *HTMLElement {
	if node == nil {
		return nil
	}
	return &HTMLElement{
		n: node,
	}
}

type TraverseFn = func(element *HTMLElement) TraverseGuide

func (el *HTMLElement) Traverse(callback TraverseFn) {
	if el == nil {
		return
	}
	// call it on the current node
	guide := callback(el)
	if guide.FirstChild {
		// traverse to the first child
		NewHTMLElement(el.n.FirstChild).Traverse(callback)
	}
	if guide.NextSibling {
		// traverse to the first sibling
		NewHTMLElement(el.n.NextSibling).Traverse(callback)
	}
}

type FindFn = func(element *HTMLElement) bool

func (el *HTMLElement) Find(callback FindFn) *HTMLElement {
	if el == nil {
		return nil
	}
	// call it on the current node
	if callback(el) {
		return el
	}
	if found := NewHTMLElement(el.n.FirstChild).Find(callback); found != nil {
		return found
	}
	if found := NewHTMLElement(el.n.NextSibling).Find(callback); found != nil {
		return found
	}
	return nil
}

func (el *HTMLElement) String() string {
	return fmt.Sprintf(
		NodeFormat,
		NodeType(el.n.Type),
		el.n.DataAtom,
		el.n.Attr,
	)
}

const NodeFormat = `{
	type: %s
	element: %s
	attr: %v
}
`

type NodeType html.NodeType

func (nt NodeType) String() string {
	switch html.NodeType(nt) {
	case html.ErrorNode:
		return "ErrorNode"
	case html.TextNode:
		return "TextNode"
	case html.DocumentNode:
		return "DocumentNode"
	case html.ElementNode:
		return "ElementNode"
	case html.CommentNode:
		return "CommentNode"
	case html.DoctypeNode:
		return "DoctypeNode"
	case html.RawNode:
		return "RawNode"
	default:
		return "unknown"
	}
}

func (el *HTMLElement) HasClasses(classNames ...string) bool {
	clsIdx := slices.IndexFunc(el.n.Attr, func(attr html.Attribute) bool { return attr.Key == "class" })
	if clsIdx == -1 {
		return false
	}
	cls := strings.Split(el.n.Attr[clsIdx].Val, " ")
	for _, name := range classNames {
		if !slices.Contains(cls, name) {
			return false
		}
	}
	return true
}

func (el *HTMLElement) Is(atomType atom.Atom) bool {
	return el.n.DataAtom == atomType
}

func (el *HTMLElement) Parent() *HTMLElement {
	return NewHTMLElement(el.n.Parent)
}

func (el *HTMLElement) FirstChild() *HTMLElement {
	return NewHTMLElement(el.n.FirstChild)
}

func (el *HTMLElement) Content() string {
	childContent := el.FirstChild().Find(func(element *HTMLElement) bool {
		return element.n.Type == html.TextNode
	})
	if childContent == nil {
		return ""
	}
	return childContent.n.Data
}

// curl -s -S 'https://www.cuballama.com/viajes/activity/search' \
//   -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' \
//   -H 'Accept-Language: en-US,en;q=0.9' \
//   -H 'Cache-Control: no-cache' \
//   -H 'Connection: keep-alive' \
//   -H 'Content-Type: application/x-www-form-urlencoded' \
//   -H 'Cookie: PHPSESSID=a4cfdf069f228ced38ebf977bc39952e; wmsession=d0d1c396-284a-439d-915a-b2dd47c864ce-1732145711047; wm_lang_code=en' \
//   -H 'DNT: 1' \
//   -H 'Origin: https://www.cuballama.com' \
//   -H 'Pragma: no-cache' \
//   -H 'Referer: https://www.cuballama.com/viajes/activity/search' \
//   -H 'Sec-Fetch-Dest: document' \
//   -H 'Sec-Fetch-Mode: navigate' \
//   -H 'Sec-Fetch-Site: same-origin' \
//   -H 'Upgrade-Insecure-Requests: 1' \
//   -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36' \
//   -H 'sec-ch-ua: "Not;A=Brand";v="24", "Chromium";v="128"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "macOS"' \
//   --data-raw 'attr_start_date=11%2F02%2F2025&attr_end_date=18%2F02%2F2025&attr_destination_iata=VRO&attr_destination_name=Matanzas'
