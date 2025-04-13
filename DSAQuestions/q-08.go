package main

import "fmt"

/*
	8. Implement merge sort.
*/

func merge(nums1, nums2 []int) []int {
	l1 := len(nums1)
	l2 := len(nums2)
	result := make([]int, l1+l2)

	i, j := 0, 0

	for i < l1 && j < l2 {
		if nums1[i] < nums2[j] {
			result[i+j] = nums1[i]
			i++
		} else {
			result[i+j] = nums2[j]
			j++
		}
	}

	for i < l1 {
		result[i+j] = nums1[i]
		i++
	}

	for j < l2 {
		result[i+j] = nums2[j]
		j++
	}
	return result

}

func main() {
	nums1 := []int{1, 2, 4, 5, 67, 88}
	nums2 := []int{34, 56, 87, 90}

	fmt.Println("result : ", merge(nums1, nums2))

}
