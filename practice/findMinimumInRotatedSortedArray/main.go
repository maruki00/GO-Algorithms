package main

import "fmt"

func findMin(nums []int) int {

	numsLen := len(nums)
	if numsLen == 1 {
		return nums[0]
	}
	l, h := 0, numsLen-1
	if nums[h] > nums[l] {
		return nums[0]
	}

	for nums[l] > nums[h] {
		h--
	}
	return nums[h+1]
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}
