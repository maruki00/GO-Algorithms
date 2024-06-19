package main

import "fmt"

func twoSums(nums []int, target int) []int {

	ans := 0
	var arrMap map[int]int
	arrMap := make(map[int]int) //, len(nums))

	for index, item := range nums {
		fmt.Print(index, item)
	}

	// if conta arrMap

	return []int{left, right}
}

func main() {
	items := []int{1, 4, 7, 9}
	x := twoSums(items, 6)
}
