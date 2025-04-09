package main

import "fmt"

func pivotIndex(nums []int) int {
	result := 0
	sumArray := 0
	sumLift := make([]int, len(nums))
	sumRight := make([]int, len(nums))
	j, i := 0, len(nums)-1
	for i := range nums {
		sumArray += nums[i]
	}
	sumRight[i] = sumArray
	sumLift[j] = sumArray
	i--
	j++
	for ; i >= 0; i-- {
		sumLift[i] -= nums[i]
	}

	for ; j < len(nums); j++ {
		sumRight[j] -= nums[j]
	}
	for i := 0; i < len(nums); i++ {
		if sumLift[i] == sumRight[i] {
			result++
		}
	}
	return result
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("result : ", pivotIndex(nums))
}
