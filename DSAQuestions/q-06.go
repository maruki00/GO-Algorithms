package main

import (
	"fmt"
)

/*
	6. Implement binary search in an array.
*/

func find(nums []int, target int) int {
	middle := 0
	first, last := 0, len(nums)-1
	for first < last {
		middle = (first + last) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			first = middle + 1
		} else {
			last = middle - 1
		}
	}
	return -1
}
func main() {
	nums := []int{1, 3, 4, 6, 7, 8, 9, 10, 13, 16}
	target := 10
	fmt.Println("result : ", find(nums, target))
}
