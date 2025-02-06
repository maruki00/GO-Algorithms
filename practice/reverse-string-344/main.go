package main

import (
	"fmt"
)

func reverseString(s []byte) {
	l, r := 0, len(s)-1

	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)

	for _, char := range s {
		fmt.Println(string(char))
	}

}
