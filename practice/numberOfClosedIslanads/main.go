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
