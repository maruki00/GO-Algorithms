package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	sts := strings.Split(s, " ")
	star, end := 0, len(sts)-1
	for star <= end {
		sts[star], sts[end] = sts[end], sts[star]
		end--
		star++
	}
	return strings.Join(sts, " ")
}

func main() {
	s := "the sky is blue"
	fmt.Println("result : ", reverseWords(s))
}
