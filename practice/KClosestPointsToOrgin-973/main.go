package main

import "math"

type Heap [][]int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(a, b []int) bool {
	x1, x2 := math.Pow(float64(a[0]), 2), math.Pow(float64(b[0]), 2)
	y1, y2 := math.Pow(float64(a[1]), 2), math.Pow(float64(b[1]), 2)
	return math.Sqrt(x1+y1) > math.Sqrt(x2+y2)
}
func (h Heap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}
func (h *Heap) Pop() []int {
	old := *h
	lst := *old[len(*h)-1]
	old = old[:len(*h)-1]
}

func kClosest(points [][]int, k int) [][]int {

	return [][]int{}
}

func main() {

}
