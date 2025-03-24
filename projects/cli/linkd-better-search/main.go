package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func buildURL(start int) string {
	baseURL := "https://www.linkedin.com/jobs/search/"

	queryParams := url.Values{
		"f_E":      {"4"},
		"f_JT":     {"F"},
		"f_TPR":    {"r604800"},
		"keywords": {"software engineer"},
		"location": {"Spain"},
		"sortBy":   {"R"},
	}

	if start > 0 {
		queryParams.Add("start", strconv.Itoa(start))
	}

	return fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
}

func main() {
	keywords := []string{}
	// includeLocations := []string{"spain"}
	blacklistTitles := []string{
		"java",
		"javascript",
		"php",
		"android",
		".net",
		"c++",
		"c#",
		"c",
		"node.js",
		"junior",
		"ingeniero",
		"programador",
		"desarrollador",
		"robot",
		"react",
		"vue",
		"angular",
		"fullstack",
		"full-stack",
		"frontend",
		"front-end",
		"front end",
	}

	// url := "https://www.linkedin.com/jobs/search/?f_E=4&f_JT=F&f_TPR=r604800&keywords=softwere%20developer&location=Madrid%2C%20Community%20of%20Madrid%2C%20Spain&sortBy=R"
	searchResult, err := SearchJobPage(0)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	jobs := searchResult.Jobs
	jobCount := searchResult.Count
	fmt.Printf("jobCount: %v\n", jobCount)
	fmt.Printf("len(jobs): %v\n", len(jobs))

	filteredJobs := Filter(jobs, FilterJobItems(keywords, blacklistTitles))

	fmt.Printf("len(filteredJobs): %v\n", len(filteredJobs))
	for index, jobInfo := range filteredJobs {
		fmt.Printf("index: %d, job: %s\n", index, jobInfo)
	}

	// fmt.Println("wait for 1sec")
	// time.Sleep(time.Second)
	// searchResult, err = SearchJobPage(len(jobs))
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	return
	// }
	// jobs = searchResult.Jobs
	// jobCount = searchResult.Count
	// fmt.Printf("jobCount: %v\n", jobCount)
	// fmt.Printf("len(jobs): %v\n", len(jobs))
	//
	// filteredJobs = Filter(jobs, FilterJobItems(keywords, blacklistTitles))
	//
	// fmt.Printf("len(filteredJobs): %v\n", len(filteredJobs))
	// for index, jobInfo := range filteredJobs {
	// 	fmt.Printf("index: %d, job: %s\n", index, jobInfo)
	// }

}

func FilterJobItems(keywords, blacklistTitles []string) func(*JobItem) bool {
	return func(job *JobItem) bool {
		for _, kword := range keywords {
			if !strings.Contains(strings.ToLower(job.Title), kword) {
				return false
			}
		}
		for _, kword := range blacklistTitles {
			if strings.Contains(strings.ToLower(job.Title), kword) {
				// fmt.Printf("job.Title: %v\n", job.Title)
				return false
			}
		}
		// for _, location := range includeLocations {
		// 	if !strings.Contains(strings.ToLower(job.Location), location) {
		// 		return false
		// 	}
		// }
		return true
	}
}

type SearchResult struct {
	Jobs  []*JobItem
	Count int
}

func SearchJobPage(start int) (*SearchResult, error) {
	url := buildURL(start)
	fmt.Printf("searching url: %s\n", url)
	doc, err := GetHtmlFromUrl(url)
	if err != nil {
		return nil, err
	}
	main_node := findNodeBy(doc, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.DataAtom == atom.Main
	})

	jobCountStr := getDirectChildContent(findNodeBy(main_node, func(n *html.Node) bool {
		return n.Type == html.ElementNode &&
			n.DataAtom == atom.Span &&
			withAttrEval(n.Attr, attrKeyVal("class", "results-context-header__job-count"))
	}))
	jobCount, err := strconv.Atoi(jobCountStr)
	if err != nil {
		return nil, err
	}

	jobList := findNodeBy(main_node, func(n *html.Node) bool {
		return n.Type == html.ElementNode &&
			n.DataAtom == atom.Ul &&
			withAttrEval(n.Attr, attrKeyVal("class", "jobs-search__results-list"))
	})

	jobs := mapElChildNodes(jobList, ParseNodeToJobItem)

	return &SearchResult{
		Jobs:  jobs,
		Count: jobCount,
	}, nil

}

func GetHtmlFromUrl(url string) (*html.Node, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// hacking
	// req.Header.Set("User-Agent", "PostmanRuntime/7.37.3")

	fmt.Printf("req.Header: %v\n", req.Header)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status %s from url: %s", resp.Status, url)
	}

	// fmt.Printf("%#v", *resp)

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing the response body: %w", err)
	}

	return doc, nil
}

func GetMainContent(doc *html.Node) (*html.Node, error) {
	return nil, nil
}

func findManyNodeBy(initialNode *html.Node, fn func(*html.Node) bool) []*html.Node {
	stack := NewStack[*html.Node]()
	stack.Push(initialNode)

	nodes := make([]*html.Node, 0)
	for stack.Len() != 0 {
		node, _ := stack.Pop()

		if fn(node) {
			nodes = append(nodes, node)
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			stack.Push(c)
		}
	}

	return nodes
}

func findNodeBy(initialNode *html.Node, fn func(*html.Node) bool) *html.Node {
	stack := NewStack[*html.Node]()
	stack.Push(initialNode)

	for stack.Len() != 0 {
		node, _ := stack.Pop()

		if fn(node) {
			return node
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			stack.Push(c)
		}
	}

	return nil
}

func withAttrEval(attr []html.Attribute, eval func(a html.Attribute) bool) bool {
	for _, a := range attr {
		if eval(a) {
			return true
		}
	}
	return false
}

func attrKey(key string) func(a html.Attribute) bool {
	return func(a html.Attribute) bool {
		return a.Key == key
	}
}

func attrKeyVal(key, val string) func(a html.Attribute) bool {
	return func(a html.Attribute) bool {
		return a.Key == key && a.Val == val
	}
}

func attrKeyPartialVal(key, val string) func(a html.Attribute) bool {
	return func(a html.Attribute) bool {
		return a.Key == key && strings.Contains(a.Val, val)
	}
}

func getDirectChildContent(node *html.Node) string {
	if node == nil {
		return ""
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode && c.Data != "" {
			return strings.TrimSpace(c.Data)
		}
	}
	return ""

}

func getAttrContent(node *html.Node, key string) string {
	if node == nil {
		fmt.Printf("get attr: %s from node: %s\n", key, &HtmlNodePrinter{node})
		return ""
	}
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func mapElChildNodes[T any](node *html.Node, mapper func(n *html.Node) T) []T {
	// fmt.Printf("node: %s\n", &HtmlNodePrinter{node})
	arr := make([]T, 0)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			// fmt.Printf("node: %s\n", &HtmlNodePrinter{c})
			arr = append(arr, mapper(c))
		}
	}
	return arr

}

type JobItem struct {
	Title       string
	Date        time.Time
	DetailsLink string
	Location    string
}

func ParseNodeToJobItem(n *html.Node) *JobItem {
	jItem := &JobItem{}
	jItem.Title = getDirectChildContent(findNodeBy(n, func(n *html.Node) bool {
		return n.Type == html.ElementNode &&
			n.DataAtom == atom.H3 &&
			withAttrEval(n.Attr, attrKeyVal("class", "base-search-card__title"))
	}))
	jItem.DetailsLink = getAttrContent(findNodeBy(n, func(n *html.Node) bool {
		return n.Type == html.ElementNode &&
			n.DataAtom == atom.A &&
			withAttrEval(n.Attr, attrKeyPartialVal("href", "https://es.linkedin.com/jobs/view/"))
	}), "href")
	jItem.Date, _ = time.Parse("2006-01-02", getAttrContent(findNodeBy(n, func(n *html.Node) bool {
		return n.Type == html.ElementNode &&
			n.DataAtom == atom.Time &&
			withAttrEval(n.Attr, attrKey("datetime"))
	}), "datetime"))
	jItem.Location = getDirectChildContent(findNodeBy(n, func(n *html.Node) bool {
		return n.Type == html.ElementNode &&
			n.DataAtom == atom.Span &&
			withAttrEval(n.Attr, attrKeyVal("class", "job-search-card__location"))
	}))

	return jItem
}

func (ji *JobItem) String() string {
	return fmt.Sprintf(
		"{title: %s, date: %s, location: %s}",
		ji.Title,
		ji.Date,
		ji.Location,
	)
}
