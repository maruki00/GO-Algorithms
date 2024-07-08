package main

import (
	"fmt"
	"sort"
)

func topKElement(nums []int, k int) []int {
	freq := make(map[int]int)
	out := []int{}
	for _, j := range nums {
		freq[j]++
	}
	keys := []int{}
	for _, j := range freq {
		keys = append(keys, j)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return freq[keys[i]] >= k && freq[keys[i]] > freq[keys[j]]
	})
	for _, j := range keys {
		if freq[j] >= k {
			fmt.Println("iotem : ", freq[j], j)
			out = append(out, j)
		}
	}
	return out
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println("result : ", topKElement(nums, 2))
}
