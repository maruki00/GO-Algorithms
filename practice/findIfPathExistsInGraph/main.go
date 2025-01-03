package main

type Stack []int

func (s Stack) pop() int {
	last := s[len(s)-1]
	s = s[:len(s)-1]
	return last
}

func (s Stack) push(item ...int) {
	s = append(s, item...)
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	visited := make(map[int]bool, 0)
	graph := make(map[int][]int, 0)
	for _, a := range edges {
		graph[a[0]] = append(graph[a[0]], a[1])
	}

	var BFS func(visited map[int]bool, node int)
	BFS = func(visited map[int]bool, node int) {

	}

	BFS(visited, 0)
	return false
}

func main() {
	validPath(1, nil, 0, 0)
}
