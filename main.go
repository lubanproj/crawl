package main

func main() {
	// environment variable preparation
	GetValueFromEnv("GITHUB_TOKEN")

	// crawl
	Crawl("https://gocn.vip/topics/node18")
	// push to github

}
