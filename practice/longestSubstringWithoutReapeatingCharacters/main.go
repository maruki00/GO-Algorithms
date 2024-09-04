package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {

	memorized := make(map[string]bool, 0)
	prev := ""
	l, h := 0, 0
	maxConter := 0
	for h < len(s) {
		if !strings.Contains(prev, string(s[h])) {
			prev += string(s[h])
			h++
			maxConter = max(maxConter, len(prev))
			continue
		}
		if _, ok := memorized[prev]; !ok {

			maxConter = max(maxConter, len(prev))

		}
		memorized[prev] = true
		prev = prev[1:]
		l++

	}
	return maxConter
}

func main() {
	s := " "
	fmt.Println("resultt : ", lengthOfLongestSubstring(s))

}
