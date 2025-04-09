package main

import "fmt"

func pivotIndex(nums []int) int {
	result, sumArray := 0, 0
	j, i := 0, len(nums)-1
	sumLift := make([]int, len(nums))
	sumRight := make([]int, len(nums))
	for i := range nums {
		sumArray += nums[i]
	}
	sumRight[i] = sumArray - nums[i]
	sumLift[j] = sumArray - nums[j]
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
	if result == 0 {
		return -1
	}
	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("result : ", pivotIndex(nums))
}
