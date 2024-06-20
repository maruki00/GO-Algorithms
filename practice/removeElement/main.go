package main

import "fmt"

func removeElement(nums []int, val int) int {

	k := len(nums)
	position := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			k--
		} else {
			nums[position] = nums[i]
			position++
		}
	}
	return k
}

func main() {
	items := []int{2, 2, 2}
	fmt.Println(removeElement(items, 2))
}
