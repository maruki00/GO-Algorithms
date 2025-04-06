package main

import (
	"fmt"
	"sort"
)

// func closeStrings(word1 string, word2 string) bool {
// 	w1Map := make(map[byte]int)
// 	w2Map := make(map[byte]int)
// 	for i := range word1 {
// 		w1Map[word1[i]]++
// 	}
// 	for i := range word2 {
// 		w2Map[word2[i]]++
// 	}
// 	for k, v := range w1Map {
// 		if ex, ok := w2Map[k]; !ok || v != ex {
// 			return false
// 		}
// 	}
// 	return true
// }

func closeStrings(word1 string, word2 string) bool {
	counter := func(word string) (keys, vals [26]int) {
		for i := range word {
			keys[word[i]-'a'] = 1
			vals[word[i]-'a'] += 1
		}
		sort.Ints(vals[:])
		return keys, vals
	}
	keys1, vals1 := counter(word1)
	keys2, vals2 := counter(word2)
	return keys1 == keys2 && vals1 == vals2
}

func main() {
	// word1 := "cabbba"
	// word2 := "abbccc"
	word1 := "abcbb"
	word2 := "abccc"
	fmt.Println("result : ", closeStrings(word1, word2))
}
