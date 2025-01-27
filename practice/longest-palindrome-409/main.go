package main

func longestPalindrome(s string) int {

	sMap := make(map[byte]int, 0)
	result := 0
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
		if sMap[s[i]]%2 == 0 {
			result += 2
			delete(sMap, s[i])
		}
	}

	if len(sMap) > 0 {
		result += 1
	}
	return result
}

func main() {

}
