package main

import "fmt"

func pivotIndex(nums []int) int {
	totalL := 0
	totalR := 0
	for i := range nums {
		totalL += nums[i]
	}

	for i, w := range nums {
		totalL -= w
		if totalR == totalL {
			return i
		}
		totalR += w
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("result : ", pivotIndex(nums))
}
