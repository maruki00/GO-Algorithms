package main

import (
	"fmt"
	"sort"
)

func topKElement(nums []int) int {
	freq := make(map[int]int)
	for _, j := range nums {
		freq[j]++
	}
	keys := []int{}
	for _, j := range freq {
		keys = append(keys, j)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return freq[keys[i]] < freq[keys[j]]
	})
	fmt.Println("result : ", freq, keys)
	return 1
}

func main() {
	_ = topKElement(1, 1, 1, 2, 2, 3)
}
