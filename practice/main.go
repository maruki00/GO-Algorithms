package main

import "fmt"

func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			i++
		}
		if nums[j] != 0 {
			j++
		}
		if nums[i] != 0 && nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
			i++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeros(nums)
	fmt.Println("result : ", nums)
}
