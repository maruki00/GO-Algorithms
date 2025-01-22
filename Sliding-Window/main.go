package main

import "fmt"

func MaxSubArray(nums []int, subLen int) int {
	result := 0
	for i := 0; i < subLen; i++ {
		result += nums[i]
	}
	for i := subLen; i < len(nums); i++ {
		result = max(result, nums[i]+(result-nums[i-1]))
	}
	return result
}

func MinSubArray(nums []int, subLen int) int {
	result := 0
	for i := 0; i < subLen; i++ {
		result += nums[i]
	}
	for i := subLen; i < len(nums); i++ {
		result = min(result, nums[i]+(result-nums[i-1]))
	}
	return result
}
func main() {
	nums := []int{4, 5, 6, 8, 3, 65, 8, 3, 5, 7, 8, 2, 4, 6}
	ln := 3
	fmt.Println("result : ", MaxSubArray(nums, ln))
	fmt.Println("result : ", MinSubArray(nums, ln))
}
