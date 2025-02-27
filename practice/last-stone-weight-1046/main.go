package main

import (
	"container/heap"
	"fmt"
)

func lastStoneWeight(stones []int) int {
	maxHeap := &MaxHeap{}

	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}

	for maxHeap.Len() > 1 {
		stone1 := heap.Pop(maxHeap).(int)
		stone2 := heap.Pop(maxHeap).(int)

		if stone1 != stone2 {
			heap.Push(maxHeap, stone1-stone2)
		}
	}

	res := 0
	if maxHeap.Len() == 1 {
		res = (*maxHeap)[0]
	}
	return res
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println("result : ", lastStoneWeight(stones))
}
