package main

import "fmt"

func majorityElement(nums []int) int {
	hashSet := make(map[int]int)
	majorELement, count := 0, 0
	for _, item := range nums {
		hashSet[item]++
		if count < hashSet[item] {
			count = hashSet[item]
			majorELement = item
		}
	}
	return majorELement
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 1}
	fmt.Println("Result : ", majorityElement(nums))
}
