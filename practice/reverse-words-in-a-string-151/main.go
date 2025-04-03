package main

import (
	"fmt"
	"strings"
)

func split(s string) []string {
	i := 0
	result := make([]string, 0)
	for i < len(s) {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		j := i
		for i < len(s) && s[i] != ' ' {
			i++
		}
		if i != j {
			result = append(result, string(s[j:i]))
		}
	}
	return result
}
func reverseWords(s string) string {
	sts := split(s)
	star, end := 0, len(sts)-1
	for star <= end {
		sts[star], sts[end] = sts[end], sts[star]
		end--
		star++
	}
	return strings.Join(sts, " ")
}

func main() {
	// s := "the sky is blue"
	s := "a good   example"
	fmt.Println("result : ", reverseWords(s))
}
