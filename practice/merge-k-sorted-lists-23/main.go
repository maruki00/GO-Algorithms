package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Heap []*ListNode

var current = 0

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(a, b int) bool { return h[a].Val > h[b].Val }
func (h Heap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]

}
func (h *Heap) Pop() interface{} {
	old := *h
	lst := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return lst
}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func mergeKLists(lists []*ListNode) *ListNode {

	h := &Heap{}
	heap.Init(h)
	for _, node := range lists {
		x := node
		for x != nil {
			heap.Push(h, x)
			x = x.Next
		}
	}
	var result *ListNode
	for h.Len() > 0 {
		x := heap.Pop(h).(*ListNode)
		result = &ListNode{Val: x.Val, Next: result}
	}

	return result
}

// [[1,4,5],[1,3,4],[2,6]]
func main() {
	lists := []*ListNode{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeKLists(lists))
}
