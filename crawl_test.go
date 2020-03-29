package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/gocolly/colly"
)

//func TestCrawl(t *testing.T) {
//	Crawl("https://gocn.vip/news")
//}


func TestColly(t *testing.T) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[title]", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		fmt.Println(e.Attr("title"), e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// c.Visit("http://go-colly.org/")
	c.Visit("https://gocn.vip/topics/node18")
}

func TestParseContent(t *testing.T) {
	title := ""
	_ , ok := parseContent(title)
	assert.Equal(t, ok, false)

	body := "<ol><li>dubbogo v1.4 新特性 <a href=\"https://gocn.vip/topics/10119\" rel=\"nofollow\" target=\"_blank\">" +
		"https://gocn.vip/topics/10119</a></li><li>golang 快速入门 [7.2]-北冥神功—go module 绝技 <a href=\"https://gocn.vip/topics/10118\" " +
		"rel=\"nofollow\" target=\"_blank\">https://gocn.vip/topics/10118</a><br></li><li>Go 每日一库之 jj" +
		" <a href=\"https://segmentfault.com/a/1190000022163724\" rel=\"nofollow\" target=\"_blank\">https://segmentfault.com/a/1190000022163724</a>" +
		"</li><li>Go 官方的限流器 time/rate 如何使用 <a href=\"https://mp.weixin.qq.com/s/QMbZOh8_LGIUdIdEBQ2-yA\" rel=\"nofollow\" target=\"_blank\">" +
		"https://mp.weixin.qq.com/s/QMbZOh8_LGIUdIdEBQ2-yA</a></li><li>分布式从 ACID、CAP、BASE 的理论推进 <a href=\"https://gocn.vip/topics/10121\" " +
		"rel=\"nofollow\" target=\"_blank\">https://gocn.vip/topics/10121</a></li></ol>"

	res , ok := parseContent(body)
	fmt.Println(res, ok)
}