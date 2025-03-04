package main

import (
	"fmt"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)

	res := nums[0] + nums[1] + nums[2]
	mn := max(res, target) - min(res, target)
	total := res
	for i := 0; i < len(nums); i++ {
		p1 := i + 1
		p2 := len(nums) - 1
		for p1 < p2 {
			if sum := nums[i] + nums[p1] + nums[p2]; sum < target {
				res = sum
				p1++
			} else if sum > target {
				res = sum
				p2--
			} else {
				return target
			}

			if max(res, target)-min(res, target) < mn {
				mn = max(res, target) - min(res, target)
				total = res
			}
		}
	}
	return total
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println("result : ", threeSumClosest(nums, target))
}
