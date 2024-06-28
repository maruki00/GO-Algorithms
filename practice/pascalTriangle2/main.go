package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{0}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	prev := []int{1, 1}
	next := []int{}
	for i := 2; i <= rowIndex; i++ {
		next = []int{1}
		for j := 1; j < i-1; j++ {
			next = append(next, prev[j-1]+prev[j])
		}
		next = append(next, 1)
		if i == rowIndex {
			break
		} else {
			prev = next
		}
	}
	return next
}

func main() {
	fmt.Println("Result : ", getRow(3))
}
