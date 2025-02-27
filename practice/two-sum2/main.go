package main

import "fmt"

func twoSum(numbers []int, target int) []int {

	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
		if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return []int{}

}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println("result : ", twoSum(nums, 9))
}
