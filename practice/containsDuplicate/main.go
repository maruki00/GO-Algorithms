package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {

	hashSet := make(map[int]int)
	for _, i := range nums {
		hashSet[i]++
		if hashSet[i] > 1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("Result : ", containsDuplicate(nums))
}
