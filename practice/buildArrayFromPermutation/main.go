package main

import "fmt"

func buildArray(nums []int) []int {

	ret := []int{}
	for _, j := range nums {
		ret = append(ret, nums[j])
	}
	return ret
}

func main() {
	nums := []int{1, 3, 5, 2, 6, 4, 0}
	fmt.Println("Result : ", buildArray(nums))

}
