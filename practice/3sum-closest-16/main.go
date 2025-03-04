package main

import "fmt"

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func threeSumClosest(nums []int, target int) int {
	resu := 0
	i := 0
	for ; i < 3 && i < len(nums); i++ {
		resu += nums[i]
	}
	closeset := resu

	for ; i < len(nums)-3; i++ {
		resu -= nums[i]
		resu += nums[i+3]
		closeset = min(closeset, abs(target-resu))
	}
	return closeset
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println("result : ", threeSumClosest(nums, target))
}
