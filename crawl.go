package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"github.com/lubanproj/gorpc/log"
	"os"
	"strings"
)

// Get the configuration from the computer environment variables
func GetValueFromEnv(key string) string {
	return os.Getenv(key)
}

// environment variable preparation
var Token = GetValueFromEnv("GITHUB_TOKEN")


// Crawl all gocn topics
func Crawl(url string) {

	pattern := `/topics/\d+`

	collector := colly.NewCollector()
	collector.OnHTML("a[title]", func(e *colly.HTMLElement) {
		// regex match topic
		path := e.Attr("href")
		topic, ok := regexMatch(path, pattern)
		if ok {
			e.Request.Visit(fmt.Sprintf("https://gocn.vip%s",topic))
		}
	})

	redisAddr := ":6379"
	conn, err := redis.Dial("tcp",redisAddr)
	if err != nil {
		log.Fatalf("get redis conn error : %v", err)
	}
	defer conn.Close()

	collector.OnRequest(func(r *colly.Request) {
		topic, ok := regexMatch(r.URL.Path, pattern)
		if ok {
			r.Visit(fmt.Sprintf("https://gocn.vip%s",topic))
			// fmt.Println("content",r.URL)
		}

	})

	collector.OnResponse(func(r *colly.Response) {
		content, ok := parseContent(string(r.Body))
		fmt.Println("content : ", content)

		if ok && content != ""{
			pushToGithub(content, Token)
		}
	})

	collector.Visit(url)
}


func parseContent(body string) (string, bool) {

	pattern := `<p>GoCN(.|\n|\t)*每日新闻(.*?)</p>`
	title, _ := regexMatch(body, pattern)
	if title == "" {
		pattern = `<h[0-9]>GoCN(.|\n|\t)*每日新闻(.|\n|\t)*</h[0-9]>?`
		title, _ = regexMatch(body, pattern)

		if title == "" {
			return "", false
		}
		pattern = `>(.|\n|\t)*每日新闻(.|\n|\t)*<`
		title, _ = regexMatch(title, pattern)
		title = strings.Replace(title, "<", "", 1)
		title = strings.Replace(title, ">", "", 1)
	}

	pattern = `<ol>(.|\n|\t)*</ol>`
	content, _ := regexMatch(body, pattern)

	return fmt.Sprintf("%s<hr>%s", title, content), true
}