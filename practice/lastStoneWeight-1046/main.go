package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {

	// lst :=

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	h := &MaxHeap{}
	heap.Init(h)

	// Add elements to the heap
	heap.Push(h, 10)
	heap.Push(h, 5)
	heap.Push(h, 20)

	// Peek the largest element
	fmt.Println("Largest element:", (*h)[0])

	// Remove elements from the heap
	fmt.Println("Popped element:", heap.Pop(h))
	heap.Push(h, 67)

	fmt.Println("Popped element:", h)
	fmt.Println("Popped element:", heap.Pop(h))
	fmt.Println("Popped element:", heap.Pop(h))
}
