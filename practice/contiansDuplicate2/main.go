package main

import (
	"fmt"
)

// func containsNearbyDuplicate(nums []int, k int) bool {

// 	hashSet := make(map[int][]int)
// 	for index, i := range nums {
// 		hashSet[i] = append(hashSet[i], index)
// 		found := false
// 		for _, j := range hashSet[i] {
// 			fmt.Println("Result : ", index, j, index-j)
// 			if index-j <= k {
// 				found = true
// 			}
// 		}
// 		if found {
// 			return true
// 		}
// 	}
// 	fmt.Println("result: ", hashSet)
// 	return false
// }

func containsNearbyDuplicate(nums []int, k int) bool {

	found := 0
	for i := 0; i > len(nums)-k; i++ {
		for j := i + k; j < len(nums); j++ {
			if nums[i] == nums[j] {
				found++
				break
			}

		}
		if found >= 2 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
