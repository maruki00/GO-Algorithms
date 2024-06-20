package main

import "fmt"

func strStr(haystack string, needle string) int {

	lenHaystacke := len(haystack)
	lenNeedle := len(needle)
	if haystack == needle {
		return 0
	}

	if lenNeedle > lenHaystacke {
		return -1
	}

	for i := 0; i < lenHaystacke; i++ {
		if haystack[i] == needle[0] && haystack[i:i+lenNeedle] == needle {
			return i
		}
	}
	return -1
}

func main() {
	string1 := "hello world"
	needl := "llo"
	fmt.Println("result : ", string1[2:4], string1[2:5] == needl)
}
