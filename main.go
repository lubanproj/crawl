package main

import "fmt"

var URL = "https://gocn.vip/topics/node18"

func main() {
	// crawl
	Crawl(URL)
	// push to github

}

// Specifies the number of pages to climb
func CrawlByPage(page int) {

	url := URL

	if page != 0 {
		url = fmt.Sprintf("%s?page=%d", url, page)
	}

	Crawl(url)
}

// Climb according to the given interval
func CrawlByInterval (start, end int) {
	for i:=0; i<=end && i>=start; i++ {
		CrawlByPage(i)
	}
}
