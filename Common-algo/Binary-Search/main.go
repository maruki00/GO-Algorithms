package main

import (
	"fmt"
)

func search(nums []int, target int) bool {
	start, end := 0, len(nums)
	for start <=
		end {
		middle := int((start + end) / 2)

		if nums[middle] == target {
			return true
		}
		if nums[middle] > target {
			end = middle - 1
			continue
		}
		start = middle + 1
	}
	return false
}

func main() {
	nums := []int{1, 4, 7, 9, 10, 12, 15, 17}
	fmt.Println("result : ", search(nums, 17))
}
