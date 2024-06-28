package main

import "fmt"

func findPalindrome(words []string) string {
	palindrome := ""
	for i := 0; i < len(words); i++ {
		l, r := 0, len(words[i])-1
		found := true
		for l < r {
			if words[i][l] != words[i][r] {
				found = false
				break
			}
			l++
			r--
		}
		if found {
			palindrome = words[i]
			break
		}
	}

	return palindrome
}

func main() {
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println("result : ", findPalindrome(words))

}
