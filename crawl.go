package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gomodule/redigo/redis"
	"github.com/lubanproj/gorpc/log"
	"os"
)

// Get the configuration from the computer environment variables
func GetValueFromEnv(key string) string {
	return os.Getenv(key)
}

var collector *colly.Collector

func init() {
	collector = colly.NewCollector()
}

// Crawl all gocn topics
func Crawl(url string) {
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	redisAddr := ":6379"
	conn, err := redis.Dial("tcp",redisAddr)
	defer conn.Close()
	if err != nil {
		log.Fatalf("get redis conn error : %v", err)
	}

	collector.OnRequest(func(r *colly.Request) {
		pattern := `/topics/\d+`
		topic, ok := regexMatch(r.URL.Path, pattern)

		if !ok {
			return
		}

		fmt.Println(topic)
		//if exist, err := existTopic(topic, conn); exist == 1 && err == nil {
		//	return
		//}
		//
		//if err = saveDB(topic, conn); err != nil {
		//	log.Errorf("save db error, %v", err)
		//}
	})

	collector.Visit(url)
}





