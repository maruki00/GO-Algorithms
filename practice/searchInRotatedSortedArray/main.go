package main

import "fmt"

func search(nums []int, target int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		middle := (l + h) / 2

		if nums[middle] == target {
			return middle
		}
		if nums[l] <= nums[middle] {
			if nums[l] <= target && target < nums[middle] {
				h = middle - 1
			} else {
				l = middle + 1
			}
		} else {
			if nums[middle] < target && target <= nums[h] {
				l = middle + 1
			} else {
				h = middle - 1
			}
		}
	}

	return -1
}
func main() {
	nums := []int{
		//4, 5, 6, 7, 0, 1, 2,
		1, 3,
	}
	target := 3

	fmt.Println(search(nums, target))
}
