package main

import "fmt"

func removeDuplicates(nums []int) int {
	index := 1
	occur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			occur++
		} else {
			occur = 1
		}

		if occur <= 2 {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func _removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[j-1] != nums[i] || nums[j-1] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	nums := []int{1, 1, 1, 3, 3, 5, 5, 5, 5, 5, 6, 6, 7, 7, 8, 9}
	fmt.Println("result : ", removeDuplicates(nums))
}
