package main

import "fmt"

func combine(n int, k int) [][]int {
	var ans [][]int
	var backtracking func(start int, temp []int)
	backtracking = func(start int, temp []int) {
		if len(temp) == k {
			comb := make([]int, len(temp))
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		for i := start; i <= n; i++ {
			backtracking(i+1, append(temp, i))
		}
	}

	backtracking(1, []int{})
	return ans
}

func main() {

	n, k := 4, 2
	fmt.Println("result : ", combine(n, k))
}
