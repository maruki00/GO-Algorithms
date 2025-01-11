package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt
	j := 0
	i := 0
	tmp := 0
	for i < len(nums) {
		fmt.Println(i, " - ", j)
		tmp += nums[j]
		if tmp >= target {
			result = min(result, j-i)
			tmp -= nums[i]
			i++
		}
		j++
		if j >= len(nums) {
			for i < len(nums) {
				if tmp >= target {
					result = min(result, j-i)
					tmp -= nums[i]
				}
				i++
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 4, 4}
	target := 7
	fmt.Println("result : ", minSubArrayLen(target, nums))
}
