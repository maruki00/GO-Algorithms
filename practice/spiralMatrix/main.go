package main

import "fmt"

func spiralOrder(matrix [][]int) []int {

	ret := []int{}
	for i := 0; i < len(matrix[0]); i++ {
		if i < len(matrix) {
			for j := 0; j < len(matrix[0]); j++ {
				fmt.Print(matrix[i][j], " - ")
			}
			break
		}
	}
	return ret
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println("result : ", spiralOrder(matrix))
}
