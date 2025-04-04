package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fName := "talesofdeamonsandgods.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("w3.thetalesofdemonsandgods.com"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./manga_scraps"),
	)

	chaptersMap := make(map[string][]string, 200)

	c.OnHTML("div.entry-content > div.separator > a[href]", func(e *colly.HTMLElement) {
		// fmt.Printf("e.Request.URL.String(): %v\n", e.Request.URL.String())
		link := e.Attr("href")
		// fmt.Printf("image: %v\n", link)
		chaptersMap[e.Request.URL.String()] = append(chaptersMap[e.Request.URL.String()], link)
	})

	// On every <a> element which has "href" attribute call callback
	c.OnHTML("div.nav-links div.nav-next > a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// fmt.Printf("next: %v\n", link)
		// // fmt.Printf("e.ChildText(\"span\"): %v\n", e.ChildText("span"))
		// if (strings.HasPrefix(link, "https://thebeginningaftertheendmanga.com/manga") ||
		// 	strings.HasPrefix(link, "https://thebeginningaftertheendmanga.com/uncategorized")) &&
		// 	e.ChildText("span") == "Next Chapter" {
		// 	// fmt.Printf("e.Attr(\"href\"): %v\n", e.Attr("href"))
		// 	// fmt.Printf("e.ChildTexts(\"span\"): %v\n", e.ChildTexts("span"))
		// 	// start scaping the page under the link found
		e.Request.Visit(link)
		// }
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting:", r.URL.String())
	})

	c.Visit("https://w3.thetalesofdemonsandgods.com/manga/tales-of-demons-and-gods-chapter-1/")

	type ch struct {
		ChapUrl string   `json:"chap-url"`
		ChapNo  string   `json:"chap-no"`
		Images  []string `json:"images"`
	}
	chapters := make([]ch, 0, len(chaptersMap))
	for nameUrl, images := range chaptersMap {
		chNo := ""
		parts := strings.Split(nameUrl, "-chapter-")
		if len(parts) > 1 {
			chNo = parts[1][:len(parts[1])-1] // the left part
		}
		parts = strings.Split(nameUrl, "-ch-")
		if len(parts) > 1 {
			chNo = parts[1][:len(parts[1])-1] // the left part
		}
		chapters = append(chapters, ch{
			ChapUrl: nameUrl,
			ChapNo:  chNo,
			Images:  images,
		})
	}

	slices.SortFunc(chapters, func(a, b ch) int {
		return valof(a.ChapNo) - valof(b.ChapNo)
	})

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(&chapters)

	fmt.Println("DONE!!")
}

func valof(s string) int {
	s = strings.TrimSpace(s)
	if s == "one" {
		return 10
	}
	if i, err := strconv.Atoi(s); err == nil {
		return i * 10
	}
	if parts := strings.Split(s, "-"); len(parts) > 1 {
		return must(strconv.Atoi(parts[0]))*10 + must(strconv.Atoi(parts[1]))
	}
	panic("WHY ?? [" + s + "]")
}

func must[T any](val T, err error) T {
	if err != nil {
		panic("cannot happend: " + err.Error())
	}
	return val
}
