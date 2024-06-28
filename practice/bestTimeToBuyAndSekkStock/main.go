package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
		if maxProfit < (prices[i] - minPrice) {
			maxProfit = prices[i] - minPrice
		}
		//minPrice = min(prices[i], minPrice)
		//maxProfit = max(maxProfit, (prices[i] - minPrice))
	}

	return maxProfit
}

func main() {
	numes := []int{2, 1, 4}
	fmt.Println("result : ", maxProfit(numes))
}
