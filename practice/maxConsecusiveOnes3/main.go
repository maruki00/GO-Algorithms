package main

import (
	"fmt"
	"math"
)

func longestOnes(nums []int, k int) int {

	maxCounter := math.MinInt

	l, h := 0, 0

	sumArray := 0
	zeros := 0
	for h < len(nums) && l < len(nums) {

		for h < len(nums) && (nums[h] == 1 && sumArray < (h-l)+1 || k > zeros) {
			if nums[h] == 0 {
				zeros++
			}
			sumArray += nums[h]
			h++
		}
		maxCounter = max(maxCounter, h-l)
		sumArray -= nums[l]
		if nums[l] == 0 {
			zeros--
		}
		l++
	}

	return maxCounter
}

func main() {
	nums := []int{
		//1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0,
		0, 0, 1, 1, 1, 0, 0,
	}

	fmt.Println("result : ", longestOnes(nums, 0))
}
