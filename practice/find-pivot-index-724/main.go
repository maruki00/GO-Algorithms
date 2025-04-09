package main

import "fmt"

func pivotIndex(nums []int) int {
	result := 0
	sumLift := make([]int, len(nums))
	sumRight := make([]int, len(nums))
	i, j := 0, len(nums)-1
	sumLift[i] = nums[i]
	sumRight[j] = nums[j]
	i++
	j--
	for i < len(nums) && j >= 0 {
		sumLift[i] = sumLift[i-1] + nums[i]
		sumRight[j] = sumRight[j+1] + nums[j]
		i++
		j--
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
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("result : ", pivotIndex(nums))
}
