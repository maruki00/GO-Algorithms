package main

import "fmt"

func isVowel(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	}
	return false
}

func maxVowels(s string, k int) int {
	maxV, tmpV := 0, 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			tmpV++
			maxV++
		}
	}
	for i := k; i < len(s)-k; i++ {
		if isVowel(s[i]) {
			tmpV--
		}
		if isVowel(s[i+k]) {
			tmpV++
		}
		maxV = max(maxV, tmpV)
	}
	return maxV
}

func main() {
	s := "abciiidef"
	k := 3
	fmt.Println(maxVowels(s, k))
}
