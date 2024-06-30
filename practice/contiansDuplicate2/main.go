package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {

	hashSet := make(map[int][]int)
	for _, i := range nums {
		fmt.Println(hashSet[i][0])
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	_ = containsNearbyDuplicate(nums, 3)
}
