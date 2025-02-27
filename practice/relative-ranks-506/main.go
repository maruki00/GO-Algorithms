package main

import (
	"container/heap"
	"fmt"
)

type Heap [][2]int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(a, b int) bool  { return h[a][0] > h[b][0] }
func (h Heap) Swap(a, b int)       { h[a], h[b] = h[b], h[a] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *Heap) Pop() interface{} {
	old := *h
	lst := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return lst
}

func findRelativeRanks(score []int) []string {
	result := make([]string, len(score))
	h := &Heap{}
	heap.Init(h)
	for index, s := range score {
		heap.Push(h, [2]int{s, index})
	}

	index := 0

	for h.Len() > 0 {
		item := heap.Pop(h).([2]int)
		switch index {
		case 0:
			result[item[1]] = "Gold Medal"
			break
		case 1:
			result[item[1]] = "Silver Medal"
			break
		case 2:
			result[item[1]] = "Bronze Medal"
			break
		default:
			result[item[1]] = fmt.Sprintf("%d", index+1)

		}
		index++
	}
	return result
}

func main() {

}
