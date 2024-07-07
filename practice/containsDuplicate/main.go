package main

import (
	"fmt"
	"sort"
)

func _containsDuplicate(nums []int) bool {

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]int)
	for i,j := range nums {
		hashMap[j]++
		if hashMap[j]>1{
			return true
		}
	}
	retunr false
}

func main() {:
	nums := []int{1, 2, 3, 1}
	fmt.Println("Result : ", containsDuplicate(nums))
}
