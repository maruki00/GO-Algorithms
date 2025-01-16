package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == right {
		return this.nums[left]
	}
	sumNum := 0
	for left < right {
		fmt.Println("item : ", left, right, sumNum)
		sumNum += this.nums[left] + this.nums[right]
		left++
		right--
	}
	return sumNum
}

func main() {
	nums := Constructor([]int{1, 4, -61, 4, -6})
	fmt.Println("result : ", nums.SumRange(0, 2))
}
