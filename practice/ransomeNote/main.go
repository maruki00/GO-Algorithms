package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	ranMap := make(map[rune]int)
	magMap := make(map[rune]int)

	for _, char := range ransomNote {
		ranMap[char]++
	}

	for _, char := range magazine {
		magMap[char]++
	}

	for key, value := range ranMap {
		if magMap[key] < value {
			return false
		}
	}
	return true
}

func main() {
	ransomNote := "ak"
	magazine := "aab"

	fmt.Println("result : ", canConstruct(ransomNote, magazine))
}
