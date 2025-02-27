package main

import (
	"fmt"
	"sort"
)

func _findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	lenNums := len(nums1)
	if lenNums%2 != 0 {
		return float64(nums1[int(lenNums/2)])
	} else {
		val := (nums1[int(lenNums/2)] + nums1[int(lenNums/2)+1]) / 2
		println(val)
		return float64(val)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	lenNums := len(nums1)
	if lenNums%2 != 0 {
		return float64(nums1[int(lenNums/2)])
	} else {
		return float64((float64(nums1[int(lenNums/2)]) + float64(nums1[int(lenNums/2)-1])) / 2)
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println("val : ", findMedianSortedArrays(nums1, nums2))
}
