package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"github.com/lubanproj/gorpc/log"
	"os"
	"strings"
	"time"
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

		topic := strings.Replace(r.Request.URL.Path,"/topics/","", -1)
		isExist, err := existTopic(conn, topic)

		// the topic has had crawled
		if isExist == 1 || err != nil {
			return
		}

		title, content, ok := parseContent(string(r.Body))
		titleAndContent := fmt.Sprintf("<h3>%s</h3>%s<hr>", title, content)
		fmt.Println("titleAndContent : ", titleAndContent)

		date := getDate(title)
		if curDay := time.Now().Format("2006-01-02"); curDay != date {
			// just climb today's data
			return
		}

		if ok && content != "" && title != "" {
			pushToGithub(titleAndContent, Token)
		}

		saveDB(conn, topic, date)
	})

	collector.Visit(url)
}


func parseContent(body string) (string, string, bool) {

	pattern := `<p>GoCN(.|\n|\t)*每日新闻(.*?)</p>`
	title, _ := regexMatch(body, pattern)
	if title == "" {
		pattern = `<h[0-9]>GoCN(.|\n|\t)*每日新闻(.|\n|\t)*</h[0-9]>?`
		title, _ = regexMatch(body, pattern)

		if title == "" {
			return "", "", false
		}
		pattern = `>(.|\n|\t)*每日新闻(.|\n|\t)*<`
		title, _ = regexMatch(title, pattern)
		title = strings.Replace(title, "<", "", 1)
		title = strings.Replace(title, ">", "", 1)
	}

	pattern = `<ol>(.|\n|\t)*</ol>`
	content, _ := regexMatch(body, pattern)

	return title, content, true
}

func getDate(title string) string {
	pattern := `[0-9]{4}-[0-1]{1}[0-9]{1}-[0-3]{1}[0-9]{1}`
	date, _ := regexMatch(title, pattern)

	return date
}