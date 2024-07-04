package main

import (
	"fmt"
)

func uniqueAppend(nums1 []int, item int) []int {
	for _, i := range nums1 {
		if i == item {
			return nums1
		}
	}
	return append(nums1, item)
}

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	//sort.Ints(nums1)
	//sort.Ints(nums2)
	hashSet := make(map[int]int)
	for _, j := range nums1 {
		hashSet[j]++
	}
	for _, i := range nums2 {
		if num, ok := hashSet[i]; ok {
			result = uniqueAppend(result, num)
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println("Result : ", intersection(nums1, nums2))
}
