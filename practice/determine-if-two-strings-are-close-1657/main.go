package main

import "fmt"

func closeStrings(word1 string, word2 string) bool {
	w1Map := make(map[byte]int)
	w2Map := make(map[byte]int)
	for i := range word1 {
		w1Map[word1[i]]++
	}
	for i := range word2 {
		w2Map[word2[i]]++
	}
	for k, v := range w1Map {
		if ex, ok := w2Map[k]; !ok || v != ex {
			return false
		}
	}
	return true
}

func main() {
	word1 := "abc"
	word2 := "bca"
	fmt.Println("result : ", closeStrings(word1, word2))
}
