package main

import "sort"

type Edge struct {
	src, dest, weight int
}

type Graph struct {
	vertices int
	edges    []Edge
}

func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}

func union(parent []int, x, y int) {
	xset := find(parent, x)
	yset := find(parent, y)
	parent[xset] = yset
}

func Kruskal(graph *Graph) []Edge {
	// Sort edges in increasing order
	sort.Slice(graph.edges, func(i, j int) bool {
		return graph.edges[i].weight < graph.edges[j].weight
	})

	parent := make([]int, graph.vertices)
	for i := range parent {
		parent[i] = -1
	}

	result := make([]Edge, 0, graph.vertices)

	for _, edge := range graph.edges {
		x := find(parent, edge.src)
		y := find(parent, edge.dest)

		if x != y {
			result = append(result, edge)
			union(parent, x, y)
		}
	}

	return result
}

func main() {

}
