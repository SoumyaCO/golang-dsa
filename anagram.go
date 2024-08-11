package main

import (
	"reflect"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := sMap[s[i]]; ok {
			sMap[s[i]] += 1
		} else {
			sMap[s[i]] = 1
		}
	}

	for j := 0; j < len(t); j++ {
		if _, ok := tMap[t[j]]; ok {
			tMap[t[j]] += 1
		} else {
			tMap[t[j]] = 1
		}
	}
	if reflect.DeepEqual(sMap, tMap) {
		return true
	} else {
		return false
	}
}
