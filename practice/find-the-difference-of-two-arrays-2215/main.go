package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map := make(map[int]bool)
	nums2Map := make(map[int]bool)
	result := make([][]int, 2)
	for _, i := range nums1 {
		nums1Map[i] = true
	}
	for _, i := range nums2 {
		nums2Map[i] = true
	}
	for k, _ := range nums2Map {
		if ext, ok := nums1Map[k]; !(ok && ext) {
			result[0] = append(result[0], k)
		}
	}
	for k, _ := range nums1Map {
		if ext, ok := nums2Map[k]; !(ok && ext) {
			result[1] = append(result[1], k)
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println("result : ", findDifference(nums1, nums2))
}
