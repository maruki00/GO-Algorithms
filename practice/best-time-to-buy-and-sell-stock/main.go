package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) int {
	bestProfit := 0
	left := 0
	for right := 0; right < len(prices); right++ {
		if prices[left] > prices[right] {
			left = right
		} else {
			bestProfit = max(bestProfit, prices[right]-prices[left])
		}
	}

	return bestProfit
}

func main() {
	numes := []int{2, 1, 4}
	fmt.Println("result : ", maxProfit(numes))
}
