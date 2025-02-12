package main

import (
	"container/list"
	"fmt"
)

var dl = []int{-1, 0, 1, 0}
var dc = []int{0, 1, 0, -1}

func lee(startX, startY int, mat [][]int) {
	n := len(mat)
	m := len(mat[0])
	queue := list.New()
	queue.PushBack([2]int{startX, startY})
	mat[startX][startY] = -1
	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		pos := e.Value.([2]int)
		x, y := pos[0], pos[1]
		for i := 0; i < 4; i++ {
			xx := x + dl[i]
			yy := y + dc[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < m && mat[xx][yy] == 0 {
				queue.PushBack([2]int{xx, yy})
				mat[xx][yy] = -1
			}
		}
	}
}

func main() {
	mat := [][]int{
		{0, 1, 0, 0},
		{0, -1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	startX, startY := 0, 0
	lee(startX, startY, mat)

	for _, row := range mat {
		fmt.Println(row)
	}
}
