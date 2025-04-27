package main

import "fmt"

func exist(board [][]byte, word string) bool {
	found := false
	m, n := len(board), len(board[0])
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if board[i][j] == '-' {
			return
		}
		if board[i][j] != word[k] {
			return
		}
		if k == len(word)-1 {
			found = true
			return
		}
		tmp := board[i][j]
		board[i][j] = '-'
		dfs(i-1, j, k+1)
		dfs(i+1, j, k+1)
		dfs(i, j-1, k+1)
		dfs(i, j+1, k+1)
		board[i][j] = tmp
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := 0
			dfs(i, j, k)
		}
	}
	return found
}

func main() {

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"

	fmt.Println("result : ", exist(board, word))
}
