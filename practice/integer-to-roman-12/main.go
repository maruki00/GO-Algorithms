package main

import (
	"fmt"
)

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for num > 0 {
		for i, nu := range nums {
			if num-nu >= 0 {
				result += syms[i]
				num -= nu
				break
			}
		}
	}
	return result
}

func main() {
	in := 3749
	fmt.Println("result : ", intToRoman(in))
}
