package main

import "fmt"

func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(t); i++ {
		if s[j] == t[i] {
			j++
		}
	}
	fmt.Println("i-j : ", j)
	return j == len(s)
}

func main() {
	// s := "gdc"
	// t := "ahbgdc"

	s := "abc"
	t := "ahbgdc"

	fmt.Println("result : ", isSubsequence(s, t))
}
