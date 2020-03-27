package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexMatch(t *testing.T) {
	res, ok := regexMatch("gocn.vip/topics/",`/topics/\d+`)
	assert.Equal(t, ok, false)
	fmt.Println(res)
	res, ok = regexMatch("gocn.vip/topics/1047",`/topics/\d+`)
	assert.Equal(t, ok, true)
	fmt.Println(res)
}
