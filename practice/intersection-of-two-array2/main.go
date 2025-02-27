package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {

	hashMap := make(map[int]int)
	result := []int{}

	for _,i := range nums1{
		hashMap[i]++
	}

	for


	return result
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println("Result : ", intersect(nums1, nums2))
}
