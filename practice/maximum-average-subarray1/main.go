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
	l, h := 0, k

	for h < len(nums) {
		prev -= float64(nums[l])
		prev += float64(nums[h])
		res = max(res, prev/float64(k))
		l++
		h++
	}
	return res
}

func main() {

	nums := []int{
		1, 12, -5, -6, 50, 3,
		//5,
	}
	fmt.Println("result : ", findMaxAverage(nums, 4))

}
