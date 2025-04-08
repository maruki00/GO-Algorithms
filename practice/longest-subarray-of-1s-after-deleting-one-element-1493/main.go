package main

import "fmt"

func longestSubarray(nums []int) int {
	// to be more effecient we should
	// optimize the code from O(n^2)
	// to O(n) using window sliding
	maxOnes := 0
	for i := range nums {
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
	}
	return maxOnes
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println("result : ", longestSubarray(nums))
}
