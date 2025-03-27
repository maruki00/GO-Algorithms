package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	result := math.MinInt

	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			result = max(result, profit)
		}
	}

	return result
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
