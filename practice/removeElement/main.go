package main

import "fmt"

func removeElement(nums []int, val int) int {

	k := len(nums)
	result := []int
	left := 0
	right := k - 1

	for left < right {
		if nums[left] == val {
			k--
			left++
		} else {
			result = append(result.nums[left])
			left++
		}

		if nums[right] == val {
			k--
			right--
		} else {
			result = append(result, nums[right])
			right--
		}
	}

	fmt.Println(result)
	return k
}

func main() {
	items := []int{3, 2, 2, 3}
	fmt.Println(removeElement(items, 3))
}
