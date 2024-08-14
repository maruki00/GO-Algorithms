package main

import "fmt"

func isSubsequence(s string, t string) bool {
	j := 0
	l := len(s)
	for i := 0; i < len(t); i++ {
		if l > j && s[j] == t[i] {
			j++
		}
	}
	return j == l
}

func main() {
	// s := "gdc"
	// t := "ahbgdc"

	s := "abc"
	t := "ahbgdc"

	fmt.Println("result : ", isSubsequence(s, t))
}
