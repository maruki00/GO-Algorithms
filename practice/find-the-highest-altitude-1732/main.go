package main

import "fmt"

func largestAltitude(gain []int) int {
	maxNumber, sumPrev := 0, 0
	for _, num := range gain {
		sumPrev += num
		maxNumber = max(maxNumber, sumPrev)
	}
	return maxNumber
}

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println("result : ", largestAltitude(gain))
}
