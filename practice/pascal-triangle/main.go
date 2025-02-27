package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	result := make([][]int, numRows)
	result[0] = []int{1}

	if numRows == 1 {
		return result
	}
	//result[1] = []int{1, 1}
	for i := 1; i < numRows; i++ {
		result[i] = append(result[i], 1)
		for j := 0; j < len(result[i-1])-1; j++ {
			result[i] = append(result[i], result[i-1][j]+result[i-1][j+1])
		}
		result[i] = append(result[i], 1)
	}
	return result
}

func main() {

	fmt.Println("result : ", generate(5))
}
