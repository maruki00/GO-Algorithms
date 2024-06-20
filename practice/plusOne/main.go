package main

import "fmt"

func plusOne(digits []int) []int {

	lenDigits := len(digits) - 1
	num := digits[lenDigits] + 1
	digits[lenDigits] = num % 10
	if num < 10 {
		return digits
	}
	tenRemain := num / 10
	for i := len(digits) - 2; i >= 0; i-- {

		num = digits[i] + tenRemain
		digits[i] = num % 10
		tenRemain = num / 10
	}
	if tenRemain != 0 {
		digits = append([]int{tenRemain}, digits...)
	}
	return digits
}

func main() {
	digits := []int{9, 9}
	fmt.Println("result : ", plusOne(digits))
}
