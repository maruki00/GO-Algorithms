package main

import "fmt"

func maxOperations(nums []int, k int) int {
	pairs := make(map[int]int)
	for _, i := range nums {
		pairs[i]++
	}
	counter := 0
	for _, num := range nums {

		nm1, ok := pairs[num]
		if nm1 <= 0 || !ok {
			continue
		}
		nm2, ok := pairs[k-num]
		if nm2 <= 0 || !ok {
			continue
		}

		if num == k-num && nm1 == 1 {
			continue
		}
		pairs[k-num]--
		pairs[num]--
		counter++
	}
	return counter
}

func main() {
	nums := []int{3, 1, 3, 4, 3}
	k := 6
	fmt.Println("result : ", maxOperations(nums, k))
}
