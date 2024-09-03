package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {

	prev := float64(0)

	for i := 0; i < k; i++ {
		prev += float64(nums[i])
	}
	res := prev / float64(k)
	l, h := 0, k+1

	for h < len(nums) {
		prev -= float64(nums[l-1])
		prev += float64(nums[h])
		res = min(res, prev/float64(k))
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
