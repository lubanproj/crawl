package main

import (
	"fmt"
	"testing"

	"github.com/gocolly/colly"
)

//func TestCrawl(t *testing.T) {
//	Crawl("https://gocn.vip/news")
//}


func TestColly(t *testing.T) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// c.Visit("http://go-colly.org/")
	c.Visit("https://gocn.vip/news")
}