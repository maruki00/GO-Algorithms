package main

import (
	"container/heap"
	"math"
)

type Heap [][]int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(a, b int) bool {
	x1, x2 := math.Pow(float64(h[a][0]), 2), math.Pow(float64(h[b][0]), 2)
	y1, y2 := math.Pow(float64(h[a][1]), 2), math.Pow(float64(h[b][1]), 2)
	return math.Sqrt(x1+y1) > math.Sqrt(x2+y2)
}
func (h Heap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *Heap) Pop() interface{} {
	old := *h
	lst := old[len(*h)-1]
	old = old[:len(*h)-1]
	return lst
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

// func (h *Heap) Push(x interface{}) { *h = append(*h, x.([]int)) }
// func (h *Heap) Pop() interface{} {
// 	old := *h
// 	lst := old[len(old)-1]
// 	old = old[:len(old)-1]
// 	*h = old
// 	return lst
// }

func kClosest(points [][]int, k int) [][]int {

	h := &Heap{}
	heap.Init(h)
	for _, item := range points {
		heap.Push(h, item)
	}
	return [][]int{}
}

func main() {

}
