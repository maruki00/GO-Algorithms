package main

import (
	"fmt"
	"sort"
)

// type MaxHeap []int

// func (h MaxHeap) Len() int {
// 	return len(h)
// }

// func (h MaxHeap) Less(i, j int) bool {
// 	return h[i] < h[j]
// }

// func (h MaxHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func (h *MaxHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *MaxHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }

// func buildHeapByInit(array []int) *MaxHeap {
// 	maxHeap := &MaxHeap{}
// 	nums := make([]int, 0)
// 	tracking := make(map[int]bool)
// 	for _, i := range array {
// 		if _, ok := tracking[i]; ok {
// 			continue
// 		}
// 		tracking[i] = true
// 		nums = append(nums, i)

// 	}
// 	fmt.Println(nums)
// 	*maxHeap = nums
// 	heap.Init(maxHeap)
// 	fmt.Println("buildHeapByInit: ", *maxHeap)
// 	return maxHeap
// }

// func thirdMax(nums []int) int {
// 	fmt.Println(nums)
// 	mHeap := buildHeapByInit(nums)

// 	ret := 0
// 	for mHeap.Len() > 0 {
// 		ret = mHeap.Pop().(int)
// 		fmt.Println(ret)
// 		if mHeap.Len() == 0 {
// 			return ret
// 		}
// 	}

//		return ret
//	}

func SetFromInts(nums []int) []int {
	tracking := make(map[int]bool)
	s := []int{}
	for _, i := range nums {
		if _, ok := tracking[i]; ok {
			continue
		}
		tracking[i] = true
		s = append(s, i)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Println(s)
	return s
}

func Len(s []int) int {
	return len(s)
}

func At(s []int, i int) int {
	if Len(s) == 0 {
		return 0
	}
	return s[i]
}

func thirdMax(nums []int) int {
	fmt.Println(nums)
	s := SetFromInts(nums)

	if Len(s) > 2 {
		return At(s, 2)
	}

	return At(s, 0)
}

func main() {
	nums := []int{2, 2, 3, 1}
	fmt.Println("result : ", thirdMax(nums))
}
