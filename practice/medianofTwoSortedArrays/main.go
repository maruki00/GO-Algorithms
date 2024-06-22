package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	lenNums := len(nums1)
	if lenNums%2 != 0 {
		return float64(lenNums)
	} else {
		return float64((lenNums / 2) + ((lenNums/2)+1)/2)
	}
}

func main() {

}
