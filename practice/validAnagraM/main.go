package main

import "fmt"

func isAnagram(s string, t string) bool {
	for i, j := range s {
		fmt.Println("result : ", i, j)
	}
	return true
}

func main() {
	fmt.Println("Result : ", isAnagram("anagram", "nagaram"))
}
