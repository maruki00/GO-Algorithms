package main

import "fmt"

func singleNumber(nums []int) int {
	res := make(map[int]int)
	minItem := 0
	for _, i := range nums {
		res[i]++
		minItem = min(minItem, res[i])
	}
	return minItem
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println("Result : ", singleNumber(nums))
}
