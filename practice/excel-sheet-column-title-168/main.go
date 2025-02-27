package main

import "fmt"

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func convertToTitle(n int) string {
	result := ""
	for n > 0 {
		fmt.Println(int((n - 1) % 26))
		result = string(letters[int((n-1)%26)]) + result
		n = (n - 1) / 26
	}
	return result
}
func main() {
	fmt.Println("result : ", convertToTitle(701))
}
