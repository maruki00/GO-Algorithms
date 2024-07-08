package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	out := []int{}
	for _, j := range nums {
		freq[j]++
	}
	keys := []int{}
	for j, _ := range freq {
		keys = append(keys, j)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})
	return out[:k]
}

func main() {
	nums := []int{-1, -1}
	fmt.Println("result : ", topKFrequent(nums, 2))
}
