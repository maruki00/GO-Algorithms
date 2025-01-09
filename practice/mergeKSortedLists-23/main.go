package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Heap  []*ListNode
func (h Heap) Len() int { return len(h)}
func (h Heap) Less(a, b int) bool { return h[a].Val < h[b].Val}
func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a]}
func (h *Heap) Pop() interface{} {
	old := *h
	//lst := old[len(old)-1]
	//old = old[:len(old)-1]
	//*h = old
	return old[len(old)-1]

}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
} 

func mergeKLists(lists []*ListNode) *ListNode {
	result := new(ListNode)

	h := &Heap{}
	heap.Init(h)
	for _, node := range lists {
		x := node
		for x != nil {
			heap.Push(h, x)
			x = x.Next
		}
	}
	return heap.Pop(h). 
}

func main() {

}
