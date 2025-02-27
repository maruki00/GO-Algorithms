package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] != nums[i]+1 {
			return nums[i] + 1
		}
	}
	if nums[0] != 1 {
		return 0
	}

	return nums[len(nums)-1] + 1
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println("Result : ", missingNumber(nums))
}
