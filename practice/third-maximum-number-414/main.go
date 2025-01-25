package main

import (
	"fmt"
	"sort"
)

type Set []int

func InitSet(nums []int) *Set {
	s := make(Set, 0)
	_set := make(map[int]bool)
	for _, n := range nums {
		_set[n] = true
	}
	for n, _ := range _set {
		s = append(s, n)
	}
	return &s
}

func thirdMax(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	if len(nums) < 2 {
		return nums[0]
	}
	set := InitSet(nums)
	old := *set
	if len(old) > 2 {
		return old[2]
	}
	return old[len(old)-1]
}
func main() {
	nums := []int{2, 2, 3, 1}
	fmt.Println(thirdMax(nums))
}
