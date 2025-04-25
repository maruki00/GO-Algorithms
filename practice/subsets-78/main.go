package main

import "fmt"

func subsets(nums []int) [][]int {

	result := make([][]int, 0)

	var backtracking func(nms []int, subnums []int)
	backtracking = func(nms []int, subnums []int) {

	}

	backtracking([]int{0}, 0)

	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("result : ", subsets(nums))
}
