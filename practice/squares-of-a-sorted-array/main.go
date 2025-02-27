package main

import (
	"fmt"
	"slices"
)

func sortedSquares(nums []int) []int {

	l, r := 0, len(nums)-1
	result := []int{}
	for l <= r {
		lSquar := nums[l] * nums[l]
		rSquar := nums[r] * nums[r]
		if lSquar < rSquar {
			result = append(result, int(rSquar))
			r--
		} else {
			result = append(result, int(lSquar))
			l++
		}
	}
	slices.Reverse(result)
	return result

}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println("", sortedSquares(nums))
}
