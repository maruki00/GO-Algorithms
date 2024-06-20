package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	x := 0
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

	for i, y := range s {
		fmt.Println(i, y)
	}

	fmt.Println(numbers, s)
	return x

}

func main() {
	romanToInt("xxii")
}
