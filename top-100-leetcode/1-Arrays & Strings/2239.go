package main

import (
	"fmt"
)

/**
Given an integer array nums of size n, return the number with the value closest to 0 in nums. If there are multiple answers, return the number with the largest value.

Example 1:
Input: nums = [-4,-2,1,4,8]
Output: 1
Explanation:
The distance from -4 to 0 is |-4| = 4.
The distance from -2 to 0 is |-2| = 2.
The distance from 1 to 0 is |1| = 1.
The distance from 4 to 0 is |4| = 4.
The distance from 8 to 0 is |8| = 8.
Thus, the closest number to 0 in the array is 1.

Example 2:
Input: nums = [2,-1,1]
Output: 1
Explanation: 1 and -1 are both the closest numbers to 0, so 1 being larger is returned.

*/

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
func findClosestNumber(nums []int) int {
	var closest = nums[0]

	for i := range nums {
		if abs(0-closest) >= abs(0-nums[i]) {
			closest = nums[i]
		}
	}
	return closest
}

func main() {
	nums := []int{-4, -2, 1, 4, 8}
	fmt.Println("result : ", findClosestNumber(nums))
}
