package main

import (
	"fmt"
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {

	hashSet := make(map[int][]int)
	for index, i := range nums {
		if len(hashSet[i]) == 0 {
			hashSet[i] = append(hashSet[i], index)
			continue
		}
		found := false
		for _, j := range hashSet[i] {
			if int(math.Abs(math.Inf(index-j))) > k {
				found = true
			}
		}
		if found {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
