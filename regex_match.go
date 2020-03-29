package main

import (
	"regexp"
)

func regexMatch(target string, regex string) (string,bool) {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return "", false
	}
	if ok := reg.MatchString(target); !ok {
		return "", false
	}

	return reg.FindString(target), true
}