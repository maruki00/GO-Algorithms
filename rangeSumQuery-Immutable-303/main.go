package main

import "fmt"

type NumArray struct {
	nums  []int
	cache []int
}

func Constructor(nums []int) NumArray {
	obj := NumArray{
		nums:  nums,
		cache: make([]int, len(nums)+1),
	}
	obj.cache[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		obj.cache[i] = obj.cache[i-1] + nums[i]
	}
	return obj
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == right {
		return this.nums[left]
	}
	return this.cache[right] - this.cache[left] + this.nums[left]
}

func main() {

	nums := Constructor([]int{1, 4, -6})
	fmt.Println(nums)
	fmt.Println("result : ", nums.SumRange(0, 2))
}
