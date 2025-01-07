package main

import (
	"container/heap"
	"fmt"
)

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(a, b int) bool  { return h[a] > h[b] }
func (h Heap) Swap(a, b int)       { h[a], h[b] = h[b], h[a] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	old := *h
	lst := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return lst
}
func findKthLargest(nums []int, k int) int {
	h := &Heap{}
	heap.Init(h)
	result := 0
	for _, num := range nums {
		heap.Push(h, num)
	}
	for range k {
		result = heap.Pop(h).(int)
	}
	return result
}

func main() {
	nums := []int{1, 2, 33, 4, 5}
	fmt.Println(findKthLargest(nums, 2))
}
