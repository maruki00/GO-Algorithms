package main

import "fmt"

func romanToInt(s string) int {
	x := 0
	var numbers map[string]int = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	fmt.Println(numbers, s)
	return x

}

func main() {
	romanToInt("xx")
}
