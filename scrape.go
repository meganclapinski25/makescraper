package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Story struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Age string `json:"age"`
	Points int `json:"points"`
	Url string `json:"url"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("tr.athing", func(e *colly.HTMLElement) {
		title := e.ChildText("span.titleline > a, a.storylink")
		url   := e.ChildAttr("span.titleline > a, a.storylink", "href")
                

		// Print link
		fmt.Printf("%s\n%s\n\n", title, url)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://news.ycombinator.com/")
}
