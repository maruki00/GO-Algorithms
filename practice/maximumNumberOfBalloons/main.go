package main

func maxNumberOfBalloons(text string) int {

	charsCount := make(map[rune]int)
	target := map[rune]int{
		'b': 1,
		'a': 1,
		'l': 2,
		'o': 2,
		'n': 1,
	}
	result := len(text)
	for _, char := range text {
		charsCount[char]++
	}

	for _, char := range "balloon" {
		result = min(result, charsCount[char]/target[char])
	}
	return result
}

func main() {

}
