package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	hashSet := make(map[int]int)

	for i, val := range nums {
		if v, ok := hashSet[val]; ok && i-v <= k {
			return true
		}
		hashSet[val] = i
	}
	return false
}

func main() {
	nums := []int{1, 0, 1, 1}
	fmt.Println(containsNearbyDuplicate(nums, 1))
}
