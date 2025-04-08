package main

import (
	"fmt"
)

func longestSubarray(nums []int) int {
	maxOnes := 0
	i := 0
	for ; i < len(nums); i++ {
		hasZero := 0
		j := i
		for ; j < len(nums); j++ {
			if nums[j] == 1 {
				continue
			}
			if nums[j] == 0 && hasZero == 1 {
				break
			}
			if nums[j] == 0 {
				hasZero = 1
				continue
			}
		}
		maxOnes = max(maxOnes, j-i-1)
		for i < len(nums) && nums[i] != 0 {
			i++
		}
	}
	return maxOnes
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println("result : ", longestSubarray(nums))
}
