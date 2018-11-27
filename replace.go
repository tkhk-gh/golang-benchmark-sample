package main

import (
	"regexp"
	"strings"
)

func replaceWithStrings(s string) string {
	return strings.Replace(s, ".", "_", -1)
}

var r = regexp.MustCompile(`.`)

func replaceWithRegex(s string) string {
	return r.ReplaceAllString(s, "_")
}
