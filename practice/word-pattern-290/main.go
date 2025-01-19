package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	sArray := strings.Split(s, " ")
	if len(pattern) != len(sArray) {
		return false
	}
	sMap := make(map[string]string)
	tMap := make(map[string]string)
	for i := 0; i < len(pattern); i++ {
		item, ok := sMap[string(pattern[i])]
		item1, ok1 := tMap[string(sArray[i])]
		if (ok || ok1) && (item != string(sArray[i])) && (string(pattern[i]) != item1) {
			return false
		}
		sMap[string(pattern[i])] = sArray[i]
		tMap[sArray[i]] = string(pattern[i])
	}
	return true
}

func main() {
	fmt.Println("result : ", wordPattern("abc", "b c a"))
}
