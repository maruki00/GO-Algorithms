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

func intersection_(nums1 []int, nums2 []int) []int {
	result := []int{}
	//sort.Ints(nums1)
	//sort.Ints(nums2)
	hashSet := make(map[int]int)
	for _, j := range nums1 {
		hashSet[j]++
	}
	for _, i := range nums2 {
		fmt.Println("item : ", i)
		if num, ok := hashSet[i]; ok {
			fmt.Println("found : ", num, ok, hashSet[i])
			result = uniqueAppend(result, i)
		}
	}
	fmt.Println("hashMap : ", hashSet)
	return result
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	freq := make(map[int]bool)

	for _, num := range nums1 {
		freq[num] = true
	}

	answer := make([]int, 0)
	for _, num := range nums2 {
		if freq[num] == true {
			freq[num] = false
			answer = append(answer, num)
		}
	}

	return answer
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println("Result : ", intersection(nums1, nums2))
}
