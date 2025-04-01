package main

import (
	"fmt"
)

func isMatch(s, t map[byte]int) bool {
	for k, v := range t {
		if v != s[k] {
			return false
		}
	}
	return true
}
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	minSubString := ""
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	for i := range t {
		tMap[t[i]]++
	}
	start, end := 0, len(t)
	for i := 0; i < len(t) && i < end; i++ {
		sMap[s[i]]++
	}

	for end < len(s) {
		for !isMatch(sMap, tMap) {
			sMap[s[end]]++
			end++
		}
		for isMatch(sMap, tMap) {
			sMap[s[start]]--
			start++
		}
		fmt.Println(string(s[start:end]))
	}
	return minSubString
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println("result : ", minWindow(s, t))
}
