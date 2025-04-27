package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {

	result := make([][]int, 0)
	combination := make([]int, 0)

	var backtracking func(index, currentSum int, combination []int)
	backtracking = func(index, currentSum int, combination []int) {

		if currentSum == target {
			result = append(result, append([]int{}, combination...))
			return
		}
		if currentSum > target {
			return
		}

		for i := index; i < len(candidates); i++ {
			combination = append(combination, candidates[i])
			backtracking(i, currentSum+candidates[i], combination)
			combination = combination[:len(combination)-1]
		}
	}
	backtracking(0, 0, combination)
	return result
}

func main() {
	fmt.Println("result : ", combinationSum([]int{1, 2, 3, 4, 5, 6}, 1))
}
