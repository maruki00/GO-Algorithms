package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	left, right := 0, 0
	retNums := []int{}

	for left < m && right < n {
		if nums1[left] < nums2[right] {
			retNums = append(retNums, nums1[left])
			left++
		} else {
			retNums = append(retNums, nums2[right])
			right++
		}
	}

	for left < m {
		retNums = append(retNums, nums1[left])
		left++
	}

	for right < n {
		retNums = append(retNums, nums2[right])
		right++
	}

	fmt.Println("result : ", retNums)
}

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)

}
