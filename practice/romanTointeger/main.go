package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	result := 0
	s = strings.ToUpper(s)
	var numbers map[string]int = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sLen := len(s) - 2

	for i := 0; i < sLen; i += 2 {
		romNum := string(s[i])
		romNum2 := string(s[i+1])
		if numbers[romNum] < numbers[romNum2] {
			result += (numbers[romNum2] - numbers[romNum])
		} else {
			result += (numbers[romNum2] + numbers[romNum])
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("xxii"))
}
