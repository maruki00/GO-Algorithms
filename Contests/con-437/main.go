package main

import "fmt"

func hasSpecialSubstring(s string, k int) bool {

	for i := 0; i <= len(s)-k; i++ {
		l, r := i, i+k
		conter := 0
		for l <= r {
			if s[l] != s[r] {
				break
			}
			l++
			r--
			conter += 2
		}
		if conter >= k {
			return true
		}
	}
	return false
}

func main() {
	s := "abc"
	k := 2
	fmt.Println("result : ", hasSpecialSubstring(s, k))
}

