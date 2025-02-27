package main

import "fmt"

func singleNumber(nums []int) int {
	res := make(map[int]int)
	minItem := nums[0]
	for _, i := range nums {

		res[i]++

		if res[minItem] > res[i] {
			minItem = i
		}
	}
	for i, j := range res {
		if j == 1 {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 0, 1, 2, 2}
	fmt.Println("Result : ", singleNumber(nums))
}
