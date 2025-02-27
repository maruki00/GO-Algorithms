package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {

	stack := []int{source}
	visited := make(map[int]bool, 0)
	graph := make(map[int][]int, 0)
	for _, a := range edges {
		graph[a[0]] = append(graph[a[0]], a[1])
		graph[a[1]] = append(graph[a[1]], a[0])
	}
	for len(stack) != 0 {
		lst := stack[len(stack)-1]
		if lst == destination {
			return true
		}
		stack = stack[:len(stack)-1]

		if !visited[lst] {
			visited[lst] = true
			stack = append(stack, graph[lst]...)
		}
	}
	return false
}

func main() {
	items := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	fmt.Println(validPath(1, items, 0, 5))
}
