package main

import "fmt"

func twoSum(nums []int, target int) []int {

	var arrMap map[int]int
	arrMap = make(map[int]int)

	for index, item := range nums {
		sub := target - item
		subIndex, found := arrMap[item]
		if found {
			return []int{subIndex, index}
		} else {
			arrMap[sub] = index
		}
	}
	return []int{-1, -1}
}

func main() {
	items := []int{1, 4, 7, 9}
	x := twoSum(items, 5)
	fmt.Print(x)
}
