package main

import "fmt"

func findCenter(edges [][]int) int {

	var hashSet map[int]int = map[int]int{}
	centerEdge := -1
	maxEdges := 0

	hashSet[edges[0][0]] += 1
	hashSet[edges[0][1]] += 1
	hashSet[edges[1][0]] += 1
	hashSet[edges[1][1]] += 1

	for _, i := range hashSet {
		if hashSet[i] > maxEdges {
			maxEdges = hashSet[i]
			centerEdge = i
		}
	}

	fmt.Println(hashSet)

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
