package main

import "fmt"

func subsets(nums []int) [][]int {

	result := make([][]int, 0)

	var backtracking func(nms []int, indx int)
	backtracking = func(nms []int, indx int) {
		if len(nms) == 0 || indx >= len(nums) {
			return
		}

		result = append(result, nms)
		backtracking(append(nms, nums[indx]), indx+1)
		// if len(nms) <= 0 {
		// 	return
		// }
		// nms = nms[:len(nms)-1]
	}

	backtracking([]int{0}, 0)

	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("result : ", subsets(nums))
}
