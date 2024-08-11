package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	ts := strings.ToLower(s)
	pattern := regexp.MustCompile(`[^A-Za-z\d]`)
	str := pattern.ReplaceAll([]byte(ts), []byte(""))
	l := 0
	r := len(str) - 1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		r -= 1
		l += 1
	}
	return true
}
