package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[i]+1 {
			return nums[i] + 1
		}
	}

	return 0
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println("Result : ", missingNumber(nums))
}
