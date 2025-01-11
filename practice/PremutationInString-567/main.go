package main

import (
	"fmt"
	"strings"
)

func equals(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	i, j := 0, len(s1)
	for i < len(s1) && j > 0 {
		if s1[i] != s2[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	l := len(s1)
	if equals(s2[:l], s2) {
		return true
	}

	for i := l; i < len(s2)-l; i++ {
		if equals(s2[i:i+l], s1) {
			return true
		}
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println("result : ", checkInclusion(s1, s2))
}
