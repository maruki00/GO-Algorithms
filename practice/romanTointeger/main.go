package main

import (
	"fmt"
)

func romanToInt(s string) int {
	result := 0
	var numbers map[string]int = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sLen := len(s) - 1
	i := 0
	for ; i < sLen; i++ {
		romNum := string(s[i])
		romNum2 := string(s[i+1])
		if numbers[romNum] < numbers[romNum2] {
			result -= numbers[romNum]
		} else {
			result += numbers[romNum]
		}
	}
	return result + numbers[string(s[i])]
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
