package main

import (
	"fmt"
	"math"
)

func minimizeArrayValue(nums []int) int {
	res, total := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		total += nums[i]
		x := float64(total / (i + 1))
		res = max(res, int(math.Ceil(x)))
	}
	return res
}
func main() {
	nums := []int{1, 3, 5, 6, 7, 3, 7}
	fmt.Println("Result : ", minimizeArrayValue(nums))
}
