package main

import "fmt"

func closedIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	visited := make(map[[2]int]bool)
	res := 0
	var dfs func(a int, b int) int
	dfs = func(a int, b int) int {
		if a == rows || b == cols || a < 0 || b < 0 {
			return 0
		}
		if grid[a][b] == 1 || visited[[2]int{a, b}] == true {
			return 1
		}
		visited[[2]int{a, b}] = true
		fmt.Println(visited)
		return min(dfs(a, b-1), dfs(a, b+1), dfs(a-1, b), dfs(a+1, b))
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 && visited[[2]int{i, j}] != true {
				res += dfs(i, j)
			}
		}
	}
	return res
}

func _closedIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1 || visited[i][j] {
			return
		}
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	var answer int
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 && !visited[i][j] {
				answer++
				dfs(i, j)
			}
		}
	}
	return answer
}

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}
	fmt.Println("Result : ", closedIsland(grid))
}
