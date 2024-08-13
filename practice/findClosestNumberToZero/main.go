package main

import (
	"fmt"
	"math"
)

func findClosestNumber(nums []int) int {
	closest := int(math.Abs(float64(nums[0])))
	for i := 1; i < len(nums); i++ {
		if 0-int(math.Abs(float64(nums[i]))) >= 0-int(math.Abs(float64(closest))) {
			closest = nums[i]
		}
	}
	return closest
}

func main() {
	// nums := []int{-4, -2, 1, 4, 8}
	nums := []int{2, -1, 1}
	fmt.Println("result : ", findClosestNumber(nums))
}
