package main

func maxProfit(prices []int) int {
	l, r := 0, len(prices)-1
	maxProfit := 0
	for l < r {

		if prices[l] > prices[r] {
			l++
		} else if prices[l] < prices[r] {
			r--
			maxProfit = max(maxProfit, prices[r]-prices[l])
		} else {
			l++
			r--
		}
	}

	return maxProfit
}

func main() {

}
