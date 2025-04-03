package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	sts := strings.Split(s, " ")
	star, end := 0, len(sts)-1
	for star <= end {
		sts[star], sts[end] = strings.TrimSpace(sts[end]), strings.TrimSpace(sts[star])
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
