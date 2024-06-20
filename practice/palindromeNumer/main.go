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

func isPalindrome(x int) bool {
	if x <= 10 {
		return false
	}
	nums := numToArray(x)
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] != nums[right] {
			return false
		}
	}

	return true

}

func main() {
	x := 123
	fmt.Println(isPalindrome(x))
}
