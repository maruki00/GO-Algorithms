package main

import "fmt"

func reverseStr(s string) (result string) {
	for _, ss := range s {
		result = string(ss) + result
	}
	return result
}

func longestPalindrome(s string) string {

	longest := string(s[0])
	if len(s) <= 1 {
		return s
	}

	if len(s) == 2 {
		if s == reverseStr(s) {
			return s
		}
		return string(s[1])
	}

	for i := 0; i < len(s); i++ {
		left, right := i, len(s)-1
		for left < right {
			if s[left] == s[right] && (right-left+2) > len(longest) {
				if s[left:right+1] == reverseStr(s[left:right+1]) {
					longest = s[left : right+1]
				}
			}
			right--
		}
	}

	return longest
}

func main() {

	ss := []string{"abcda", "aacabdkacaa", "dabad", "ac", "a", "bb"}
	for _, s := range ss {
		fmt.Println(longestPalindrome(s))
	}
}
