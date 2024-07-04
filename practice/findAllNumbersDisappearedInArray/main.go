package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	hashMap := make(map[int]bool)
	maxItem := len(nums)
	for _, j := range nums {
		hashMap[j] = true
		maxItem = max(maxItem, j)
	}
	for i := 1; i <= maxItem; i++ {
		if hashMap[i] != true {
			hashMap[i] = false
		}
	}
	nums = []int{}
	for i := 1; i <= maxItem; i++ {
		if hashMap[i] == false {
			nums = append(nums, i)
		}
	}

	return nums
}

func main() {
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println("Result : ", findDisappearedNumbers(nums1))
}
