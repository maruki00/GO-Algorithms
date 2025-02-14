package main

import "fmt"

func Sort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

func main() {
	nums := []int{1, 5, 3, 6, 3, 8, 3, -1, 7, 3, 7}
	Sort(nums)
	fmt.Println("result : ", nums)
}
