package main

import "fmt"

func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 && nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		if nums[j] != 0 {
			j++
		}
		i++
	}
}

func main() {
	nums := []int{1, 0}
	moveZeroes(nums)
	fmt.Println("result : ", nums)
}
