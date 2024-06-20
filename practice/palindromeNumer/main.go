package main

import "fmt"

func numToArray(num int) []int {
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num = (int)(num / 10)
	}
	return arr
}

func reverseNum(num int) int {
	reversed := 0
	for num > 0 {
		tmp := num % 10
		reversed = reversed*10 + tmp
		num = (int)(num / 10)
	}
	return reversed
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	println(reverseNum(x))
	return reverseNum(x) == x

}

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
