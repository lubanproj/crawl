package main

import (
	"fmt"
	"regexp"
)

func regexMatch(url string, regex string) (string,bool) {
	fmt.Println("url:", url)
	reg, err := regexp.Compile(regex)
	if err != nil {
		return "", false
	}
	if ok := reg.MatchString(url); !ok {
		return "", false
	}

	return reg.FindString(url), true
}