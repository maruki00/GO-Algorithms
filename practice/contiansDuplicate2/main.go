package main

import (
	"fmt"
	"math"
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

	found := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && int(math.Abs(float64(i-j))) <= k {
				found = true
				break
			}

		}
		if found {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 0, 1, 1}
	fmt.Println(containsNearbyDuplicate(nums, 1))
}
