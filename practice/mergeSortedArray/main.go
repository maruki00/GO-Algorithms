package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
	// left, right := 0, 0
	// var tmp []int = make([]int, n+m)
	// copy(tmp, nums1)

	// for left < m && right < n {
	// 	if tmp[left] < nums2[right] {
	// 		nums1[left+right] = tmp[left]
	// 		left++
	// 	} else {
	// 		nums1[left+right] = nums2[right]
	// 		right++
	// 	}
	// 	fmt.Println(tmp, nums1)
	// }

	// nums1 = append(nums1, tmp[left:m]...)
	// nums1 = append(nums1, nums2[right:n]...)

	// for left < m {
	// 	nums1[left+right] = tmp[left]
	// 	left++
	// }

	// for right < n {
	// 	nums1[left+right] = nums2[right]
	// 	right++
	// }

}

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}
