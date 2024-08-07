package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	perm := 0
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 1 {
				perm += 4
				if i > 0 && grid[i-1][j] == 1 {
					perm -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					perm -= 2
				}
			}
		}
	}
	return perm
}

func main() {
	island := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println("result : ", islandPerimeter(island))
}
