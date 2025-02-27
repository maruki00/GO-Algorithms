package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	l := len(word1)
	r := len(word2)
	result := ""
	i, j := 0, 0
	for i < l && j < r {
		result += string(word1[i]) + string(word2[j])
		i++
		j++
	}
	for i < l {
		result += string(word1[i])
		i++
	}
	for j < r {
		result += string(word2[j])
		j++
	}
	return result

}

func main() {
	word1 := "abc"
	word2 := "pqr"
	fmt.Println("result : ", mergeAlternately(word1, word2))
}
