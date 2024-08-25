package main

import "fmt"

func rotate(matrix [][]int) {
	matrixSize := len(matrix)

	for i := 0; i < matrixSize; i++ {
		for j := i + 1; j < matrixSize; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize/2; j++ {
			matrix[i][j], matrix[i][matrixSize-1-j] = matrix[i][matrixSize-1-j], matrix[i][j]
		}
	}

}

func main() {
	matrix := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	fmt.Println("result : ", matrix)
	rotate(matrix)
	fmt.Println("result : ", matrix)
}
