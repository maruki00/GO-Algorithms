package main

import "fmt"

func search(nums []int, target int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		middle := (l + h) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			if nums[h] > nums[middle] {
				l = middle + 1
				continue
			}
			h = middle - 1

		} else {
			if nums[h] < nums[middle] {
				l = middle - 1
				continue
			}
			l = middle + 1
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
