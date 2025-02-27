package main

import (
	"fmt"
)

// func getAncestors(n int, edges [][]int) [][]int {
// 	nodes := make([][]int, n)
// 	for _, node := range edges {
// 		nodes[node[1]] = append(nodes[node[1]], node[0])
// 	}
// 	return nodes
// }

func getAncestors(n int, edges [][]int) [][]int {
	ans := make([][]int, n)

	graph := make([][]int, n)
	for i := range graph {
		graph[i] = []int{}
		ans[i] = []int{}
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	for i := range n {
		traverse(graph, ans, i)
	}

	return ans
}

func traverse(graph [][]int, ans [][]int, last int) {
	vis := make(map[int]bool)
	vis[last] = true

	q := []int{}
	q = append(q, last)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, node := range graph[curr] {
			if !vis[node] {
				q = append(q, node)
				ans[node] = append(ans[node], last)

				vis[node] = true
			}
		}
	}
}

func main() {
	edgeList := [][]int{
		{0, 3},
		{0, 4},
		{1, 3},
		{2, 4},
		{2, 7},
		{3, 5},
		{3, 6},
		{3, 7},
		{4, 6},
	}
	fmt.Println(getAncestors(8, edgeList))
}
