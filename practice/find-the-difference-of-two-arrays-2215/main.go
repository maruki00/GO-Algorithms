package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map := make(map[int]int)
	nums2Map := make(map[int]int)
	result := make([][]int, 2)
	for _, i := range nums1 {
		nums1Map[i]++
	}
	for _, i := range nums2 {
		nums2Map[i]++
	}
	for k, v := range nums2 {
		if ext, ok := nums1Map[k]; !ok || ext != v {
			result[0] = append(result[0], v)
		}
	}
	for k, v := range nums1 {
		if ext, ok := nums2Map[k]; !ok || ext != v {
			result[1] = append(result[1], v)
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println("result : ", findDifference(nums1, nums2))
}
