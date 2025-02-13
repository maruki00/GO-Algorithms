package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+2)
	dp[len(cost)-1] = 0
	for i := len(cost) - 1; i >= 0; i-- {
		dp[i] = cost[i] + min(dp[i+1], dp[i+2])
	}
	return min(dp[0], dp[1])
}

func main() {
	nums := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(nums))
}
