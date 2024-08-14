package main

import "fmt"

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	prefix := 1
	for i := 0; i < length; i++ {
		res[i] = prefix
		prefix *= nums[i]
	}
	//fmt.Println("prefix : ", res)
	postfix := 1

	for i := length - 1; i > -1; i-- {
		res[i] *= postfix
		fmt.Println("suffix : ", postfix, res[i], nums[i])
		postfix *= nums[i]

	}

	//fmt.Println("suffix : ", res)
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println("result : ", productExceptSelf(nums))
}
