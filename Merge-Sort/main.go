package main

import "fmt"

func Sort(nums1, nums2 []int) []int {

	result := []int{}
	l1, l2 := len(nums1), len(nums2)
	ll1, ll2 := 0, 0

	for ll1 < l1 && ll2 < l2 {
		if nums1[l1] <= nums2[l2] {
			result = append(result, nums1[ll1])
			ll1++
			continue
		}
		result = append(result, nums2[ll1])
		ll2++
	}
	for ll1 < l2 {
		result = append(result, nums1[ll1])
		ll1++
	}
	for ll2 < l2 {
		result = append(result, nums2[ll2])
		ll2++
	}
	return result
}

func main() {

	nums1 := []int{1, 2, 4, 5, 5, 7, 8, 10}
	nums2 := []int{-1, 3, 5, 7, 8, 10, 11, 15}
	fmt.Println("result : ", Sort(nums1, nums2))

}
