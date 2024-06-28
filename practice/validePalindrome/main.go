package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome1(s string) bool {

	s = strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)
	l, r := 0, len(s)-1
	for l < r {

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func isPalindrome(s string) bool {

	res := ""
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			res += strings.ToLower(string(r))
		}
	}

	l, r := 0, len(res)-1
	for l < r {

		if res[l] != res[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func main() {
	s := "0 P"
	fmt.Println("result : ", isPalindrome(s))
}
