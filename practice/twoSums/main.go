package main

import "fmt"

func getIndex(items map[int]int, target int) int {
	for i, y := range items {
		if target == y {
			return i
		}
	}
	return -1
}

func twoSum(nums []int, target int) []int {

	var arrMap map[int]int
	arrMap = make(map[int]int) //, len(nums))

	for index, item := range nums {
		sub := target - item
		subIndex := getIndex(arrMap, item)
		if subIndex != -1 {
			return []int{subIndex, index}
		} else {
			arrMap[index] = sub
		}
	}

	return []int{-1, -1}
}

func main() {
	items := []int{1, 4, 7, 9}
	x := twoSum(items, 5)
	fmt.Print(x)
}
