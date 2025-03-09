package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}

	sort.Slice(nums, func(a, b int) bool {
		return nums[a] > nums[b]
	})

	fmt.Println(nums)
	return result
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println("result : ", fourSum(nums, target))
}
