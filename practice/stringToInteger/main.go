package main

import "fmt"

func myAtoi(s string) int {
	result := 0

	for _, char := range s {
		result = (result * 10) + int(char-'0')
	}
	return result
}

func main() {
	fmt.Println(myAtoi("1234"))
}
