package main

import "fmt"

func findSubstring(s string, words []string) []int {
	wordsSet := make(map[byte]int)
	sSet := make(map[byte]int)
	for _, w := range words {
		for i := range w {
			wordsSet[w[i]]++
		}
	}
	lenght := len(words) * len(words[0])
	for i := range s[:lenght] {
		sSet[s[i]]++
	}
	fmt.Println(lenght, sSet, s[:lenght])
	return []int{}
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}

	fmt.Println("result : ", findSubstring(s, words))

}
