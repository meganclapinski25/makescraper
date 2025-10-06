package main

import (
	"fmt"
	"encoding/json"
	"github.com/gocolly/colly"
	"os"
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

// create a slice for json 
	stories := make([]Story, 0, 64)
	// On every a element which has href attribute call callback
	c.OnHTML("tr.athing", func(e *colly.HTMLElement) {
		id := e.Attr("id")
		title := e.ChildText("span.titleline > a, a.storylink")
		url   := e.ChildAttr("span.titleline > a, a.storylink", "href")
        
		meta := e.DOM.Next()

		author := meta.Find("td.subtext a.hnuser").Text()
		age    := meta.Find("td.subtext span.age > a").Text()
		points := meta.Find("td.subtext span.score").Text()
		// Print link
		fmt.Printf(
			"Title: %s\nURL: %s\nAuthor: %s\nAge: %s\nPoints: %s\nID: %s\n\n",
			title, url, author, age, points, id,
		)

		// adding struct to slice for json
		stories = append(stories, Story{
			Title:  title,
			Author: author,
			Age:    age,
			Points: pts,
			Url:    url,
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://news.ycombinator.com/")

	b, _ := json.MarshalIndent(stories, "", "  ") // pretty JSON
	fmt.Println(string(b))                        // print to stdout
	_ = os.WriteFile("output.json", b, 0644)  
}
