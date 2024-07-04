package main

import "fmt"

func isAnagram(s string, t string) bool {
	hashMap := make(map[rune]bool)
	for _, j := range s {
		hashMap[j] = true
	}
	for _, j := range t {
		hashMap[j] = false
	}
	for _, j := range hashMap {
		if j == true {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Result : ", isAnagram("anagram", "nagaram"))
}
