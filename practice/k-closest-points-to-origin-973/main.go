package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Point struct {
	X, Y     int
	Distance float64
}
type Heap []Point

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(a, b int) bool {
	return h[a].Distance < h[b].Distance
}

func (h Heap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h *Heap) Pop() interface{} {
	old := *h
	lst := old[len(*h)-1]
	old = old[:len(*h)-1]
	*h = old
	return lst
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func kClosest(points [][]int, k int) [][]int {
	result := make([][]int, k)
	h := &Heap{}
	heap.Init(h)
	for _, item := range points {
		heap.Push(h, Point{X: item[0], Y: item[1], Distance: math.Sqrt(float64(item[0]*item[0] + item[1]*item[1]))})
	}
	for i := range k {
		x := heap.Pop(h).(Point)
		result[i] = []int{x.X, x.Y}
	}

	return result
}

func main() {
	nums := [][]int{{1, 3}, {-2, 2}, {1, 3}, {-2, 2}}
	fmt.Println("result : ", kClosest(nums, 2))
}
