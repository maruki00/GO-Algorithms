package main

import "fmt"

func KAlgo(nums []int) int {
	lMax := nums[0]
	gMax := nums[0]

	for i, _ := range nums {
		lMax = max(nums[i], lMax+nums[i])
		if lMax > gMax {
			gMax = lMax
		}
	}
	return gMax
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("result : ", KAlgo(nums))
}
