package main

import "fmt"

func spiralOrder(matrix [][]int) []int {

	res := []int{}
	move := "right"
	i, j := 0, 0
	k := 0
	level := 0
	mx := len(matrix) * len(matrix[0])
	for k < mx {
		switch move {
		case "right":
			{
				i, j = level, level
				for ; j < len(matrix[0])-level && k < mx; j++ {
					res = append(res, matrix[i][j])
					k++
				}
				j -= 1
				i++
				move = "down"
			}
			break

		case "down":
			{
				for ; i < len(matrix)-level && k < mx; i++ {
					res = append(res, matrix[i][j])
					k++
				}
				i -= 1
				j--
				move = "left"
			}
			break

		case "left":
			{
				for ; j >= 0+level && k < mx; j-- {
					res = append(res, matrix[i][j])
					k++
				}
				j += 1
				i--
				move = "up"
			}
			break

		case "up":
			{

				level++
				for ; i >= 0+level && k < mx; i-- {
					res = append(res, matrix[i][j])
					k++
				}
				move = "right"
			}
			break

		}

	}
	return res
}

func main() {
	// matrix1 := [][]int{
	// 	{1, 2, 3, 4},
	// 	{10, 11, 12, 5},
	// 	{9, 8, 7, 6},
	// }
	matrix := [][]int{
		{1, 2, 3, 4},
		{16, 17, 18, 5},
		{15, 24, 19, 6},
		{14, 23, 20, 7},
		{13, 22, 21, 8},
		{12, 11, 10, 9},
	}
	fmt.Println("result : ", spiralOrder(matrix))
}
