package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {

	right := len(nums)
	left := 0
	middle := (right + left) / 2
	for left < right {
		if target == nums[middle] {
			break
		}
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		}
		middle = (right + left) / 2
	}
	return middle
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println("result: ", searchInsert(nums, 5))
}
