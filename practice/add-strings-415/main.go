package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	n1, n2 := int64(0), int64(0)
	for _, nn1 := range num1 {
		n1 = int64((n1 * 10) + int64(nn1-'0'))
	}
	for _, nn1 := range num2 {
		n2 = int64((n2 * 10) + int64(nn1-'0'))
	}
	return fmt.Sprintf("%d", n1+n2)
}

func main() {
	fmt.Println(addStrings("123", "11"))
}
