package main

import "fmt"

func findCenter(edges [][]int) int {

	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}

func main() {
	nums := [][]int{
		{1, 10},
		{10, 2},
		{3, 10},
		{10, 4},
		{10, 5},
		{10, 6},
		{10, 7},
		{8, 10},
		{9, 10},
		{10, 11},
		{12, 10},

		// {1, 2},
		// {5, 1},
		// {1, 3},
		// {1, 4},
	}
	fmt.Println(findCenter(nums))
}

//  "10","6","9","3","+","-11","*","/","*","17","+","5","+"
// ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
