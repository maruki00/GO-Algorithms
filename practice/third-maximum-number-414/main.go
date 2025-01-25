package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	prev := nums[0]
	index := 1
	if len(nums) <= 2 {
		return nums[len(nums)-1]
	}
	for i := 0; i < len(nums); i++ {
		if prev == nums[i] {
			continue
		}
		prev = nums[i-1]
		if index >= 3 {
			return nums[i]
		}
		index++
	}
	return nums[len(nums)-1]
}
func main() {
	nums := []int{2, 2, 3, 1}
	fmt.Println(thirdMax(nums))
}
