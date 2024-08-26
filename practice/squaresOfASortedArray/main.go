package main

import (
	"fmt"
	"math"
)

func sortedSquares(nums []int) []int {

	l, r := 0, len(nums)-1
	result := []int{}

	for l < r {
		lSquar := math.Sqrt(float64(nums[l]))
		rSquar := math.Sqrt(float64(nums[r]))
		if lSquar < rSquar {
			result = append(result, int(lSquar))
			l++
		} else {
			result = append(result, int(rSquar))
			r--
		}
	}

	return result

}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println("", sortedSquares(nums))
}
