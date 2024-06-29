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
	fmt.Println(res)
	return minItem
}

func main() {
	nums := []int{1, 0, 1}
	fmt.Println("Result : ", singleNumber(nums))
}
