package main

import "fmt"

func search(nums []int, target int) int {
	ln := len(nums) - 1

	low, height := 0, ln

	for low < height {
		middle := (low + height) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[low] > target {
			height = middle - 1
		} else {
			low = middle + 1
		}

	}

	return -1

}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}

	fmt.Println("result : ", search(nums, 9))
}
