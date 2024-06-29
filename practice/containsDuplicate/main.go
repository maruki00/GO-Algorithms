package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("Result : ", containsDuplicate(nums))
}
