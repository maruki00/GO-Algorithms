package main

import "fmt"

func findSubstring(s string, words []string) []int {
	wordsSet := make(map[string]int)

	result := make([]int, 0)
	for _, w := range words {
		wordsSet[w]++
	}
	lenght := len(words) * len(words[0])
	lW0 := len(words[0])
	for i := 0; i < len(s)-lenght+1; i++ {
		sSet := make(map[string]int, len(words))
		for j := i; j < i+lenght && j < len(s); j += lW0 {
			sSet[s[j:j+lW0]]++
		}
		isOk := true
		for k, w := range sSet {
			if exists, ok := wordsSet[k]; !ok || w != exists {
				isOk = false
				break
			}
		}
		if isOk {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	// s := "barfoothefoobarman"
	// words := []string{"foo", "bar"}
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}

	fmt.Println("result : ", findSubstring(s, words))

}
