package main

import "fmt"

func max(set map[int]int, key int, value int) int {
	if set[key] > value {
		return key
	}
	return value
}
func findCenter(edges [][]int) int {

	var hashSet map[int]int = map[int]int{}
	centerEdge := -1
	maxEdges := 0
	for _, i := range edges {
		for _, j := range i {
			hashSet[j] += 1
			if hashSet[j] > maxEdges {
				maxEdges = hashSet[j]
				centerEdge = j
			}
		}
	}

	return centerEdge
}

func main() {
	nums := [][]int{
		{1, 2},
		{5, 1},
		{1, 3},
		{1, 4},
	}
	fmt.Println(findCenter(nums))
}
