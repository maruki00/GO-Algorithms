package main

import "fmt"

func Sort(nums []int) {

	for i := 0; i < len(nums); i++ {
		min_index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				min_index = j
			}
		}
		nums[min_index], nums[i] = nums[i], nums[min_index]
	}
}

func main() {
	nums := []int{1, 5, 7, 3, -1, 5, 0, 102}
	Sort(nums)
	fmt.Println("result : ", nums)
}
