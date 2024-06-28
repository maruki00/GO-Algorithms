package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	res := ""
	for _, i := range s {
		if i >= 97 && i <= 122 {
			res += string(i)
		}
	}
	if len(res) == 1 {
		return false
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
	s := " "
	fmt.Println("result : ", isPalindrome(s))
}
