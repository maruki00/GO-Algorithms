package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	res := float64(math.MinInt64)

	prev := float64(0)

	for i := 0; i < k; i++ {
		prev += float64(nums[i])
	}

	for i := 0; i < len(nums); i++ {
		if i+k > len(nums) {
			break
		}
		tmp := float64(0)
		for j := 0; j < k; j++ {
			tmp += float64(nums[i+j])
		}
		res = max(res, tmp/float64(k))

	}

	return res
}

func main() {

	nums := []int{
		//1, 12, -5, -6, 50, 3,
		5,
	}
	fmt.Println("result : ", findMaxAverage(nums, 1))

}
