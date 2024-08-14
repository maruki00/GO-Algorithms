package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	res := []string{}

	for i := 0; i < len(nums); i++ {
		fast := i + 1
		slow := i
		for fast < len(nums) && nums[fast] == nums[slow]+1 {
			fast++
			slow++
		}
		if i-slow == 0 {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[slow]))
			i = slow
		}
	}

	return res
}

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println("result : ", summaryRanges(nums))
}
