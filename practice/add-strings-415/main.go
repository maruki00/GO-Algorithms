package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	n1, n2 := 0, 0
	for _, nn1 := range num1 {
		n1 = (n1 * 10) + int(nn1-'0')
	}
	for _, nn1 := range num2 {
		n2 = (n2 * 10) + int(nn1-'0')
	}
	return "123"
}

func main() {
	fmt.Println(addStrings("123", "11"))
}
