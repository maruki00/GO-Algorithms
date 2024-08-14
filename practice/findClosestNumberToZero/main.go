package main

import (
	"fmt"
)

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func findClosestNumber(nums []int) int {
	closest := abs(nums[0])
	index := 0
	for i := 1; i < len(nums); i++ {
		if abs(nums[index]) == abs(nums[i]) && nums[index] > nums[i] {
			continue
		}

		if abs(0-abs(nums[i])) <= closest {

			index = i
			closest = abs(nums[i])
		}
	}
	return nums[index]
}

func main() {
	// nums := []int{-4, -2, 1, 4, 8}
	nums := [][]int{
		{2, -1, 1},
		{-4, -2, 1, 4, 8},
		{2, -1, 1},
	}

	for _, num := range nums {
		fmt.Println("result : ", findClosestNumber(num))
	}
}
