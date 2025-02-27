package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	index := 0
	indexes := make(map[int]int)
	maxIndex := 1
	indexes[index]++

	for i := 1; i < len(nums); i++ {

		if nums[i-1] == nums[i] {
			continue
		}
		if nums[i-1]+1 != nums[i] {
			index++
		}
		indexes[index]++
		maxIndex = max(maxIndex, indexes[index])

	}

	return maxIndex

}

func main() {

	nums := []int{
		//100, 4, 200, 1, 3, 2,
		0, 3, 7, 2, 5, 8, 4, 6, 0, 1,
	}
	fmt.Println("result : ", longestConsecutive(nums))
}
