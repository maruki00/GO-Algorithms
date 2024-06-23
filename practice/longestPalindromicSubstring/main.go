package main

import "fmt"

func reverseStr(s string) (result string) {
	for _, ss := range s {
		result = string(ss) + result
	}
	return result
}

func longestPalindrome(s string) string {

	longest := ""
	if len(s) <= 1 {
		return s
	}
	if len(s) == 2 {
		return string(s[0])
	}
	for i := 0; i < len(s); i++ {
		left, right := i, len(s)-1
		fmt.Println("Level 1 : ", i)
		for left < right {
			fmt.Println("Level 2 : ", left, right)
			if s[left] == s[right] && (right-left) > len(longest) {
				fmt.Println("Found ....", s[left:right+1], reverseStr(s[left:right+1]))
				if s[left:right+1] == reverseStr(s[left:right+1]) {
					longest = s[left : right+1]
				}
			}
			right--
		}
		left++
	}

	return longest
}

func main() {

	ss := []string{"dabad", "ac", "a"}
	for _, s := range ss {
		fmt.Println(longestPalindrome(s))
	}
}
