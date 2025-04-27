package main

import "fmt"

type Island struct {
	grid [][]int
	w    int
	h    int
}

func (isl *Island) dfs(i, j int) int {
	if i >= 0 && j >= 0 && i < isl.w && j < isl.h && isl.grid[i][j] == 1 {
		isl.grid[i][j] = 0
		return 1 + isl.dfs(i-1, j) + isl.dfs(i+1, j) + isl.dfs(i, j-1) + isl.dfs(i, j+1)
	}
	return 0
}

func maxAreaOfIsland(grid [][]int) int {
	best, w, h := 0, len(grid), len(grid[0])
	isl := &Island{grid, w, h}
	for i := range grid {
		for j := range grid[i] {
			if tmp := isl.dfs(i, j); tmp > best {
				best = tmp
			}
		}
	}
	return best
}

func main() {

	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	fmt.Println("result : ", maxAreaOfIsland(grid))

}
