package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	l := len(word1)
	r := len(word2)
	var result strings.Builder
	i, j := 0, 0
	for i < l && j < r {
		result.WriteByte(word1[i])
		result.WriteByte(word2[j])
		i++
		j++
	}
	for i < l {
		result.WriteByte(word1[i])
		i++
	}
	for j < r {
		result.WriteByte(word2[j])
		j++
	}
	return result.String()

}

func main() {
	word1 := "abc"
	word2 := "pqr"
	fmt.Println("result : ", mergeAlternately(word1, word2))
}
