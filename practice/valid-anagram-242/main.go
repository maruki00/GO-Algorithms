package main

import (
	"fmt"
)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[rune]int)
	for _, j := range s {
		hashMap[j]++
	}
	for _, j := range t {
		hashMap[j]--
	}
	for _, j := range hashMap {
		if j != 0 {
			return false
		}
	}
	return true
}

func _isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[byte]int)
	for j := range len(s) {
		hashMap[s[j]]++
		hashMap[t[j]]--
	}
	for _, j := range hashMap {
		if j != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Result : ", isAnagram("anagraml", "nagaramk"))
}
