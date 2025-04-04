package main

import "fmt"

func increasingTriplet(nums []int) bool {
	for j := 0; j < len(nums)-2; j++ {
		if nums[j] <= nums[j+1] && nums[j+1] <= nums[j+2] {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{20, 100, 10, 12, 5, 13}
	fmt.Println("result : ", increasingTriplet(nums))
}
