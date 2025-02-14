package main

const (
	white = 0
	black = 1
	red   = 2
	green = 3
)

func floodFill(arr [][]int, turn, row, col int) {
	if row >= len(arr) || row < 0 || col >= len(arr[0]) || col < 0 {
		return
	}

	if turn != black {
		return
	}
	if arr[row][col] != white {
		return
	}
	arr[row][col] = 0
	floodFill(arr, turn, row-1, col)
	floodFill(arr, turn, row+1, col)
	floodFill(arr, turn, row, col-1)
	floodFill(arr, turn, row, col+1)
}
func floodFillAll(arr [][]int, turn int) {
	for row := 0; row < 19; row++ {
		for col := 0; col < 19; col++ {
			floodFill(arr, turn, col, row)
		}
	}
}

func main() {
	arr := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{1, 4, 6, 3, 6, 7, 4},
		{1, 4, 6, 3, 6, 7, 4},
		{1, 4, 6, 3, 6, 7, 4},
		{1, 4, 6, 3, 6, 7, 4},
	}
	floodFillAll(arr, 0)
}
