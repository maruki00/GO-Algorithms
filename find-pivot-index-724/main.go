package main

import "fmt"

func pivotIndex(nums []int) int {
	result := 0
	sumLift := make([]int, len(nums))
	sumRight := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sumLift[i] = nums[i]
		sumRight[i] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if sumLift[i] == sumRight[i] {
			result++
		}
	}
	if result == 0 {
		return -1
	}
	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("result : ", pivotIndex(nums))
}
