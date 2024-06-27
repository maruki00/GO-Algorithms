package main

import "fmt"

func dailyTemmeratures(temperatures []int) []int {

	result := []int{}
	l := len(temperatures)

	for index := 0; index < l; index++ {
		found := index
		for j := index + 1; j < l; j++ {
			if temperatures[index] < temperatures[j] {
				found = j
				break
			}
		}
		result = append(result, found-index)
	}

	return result
}

func main() {

	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println("result : ", dailyTemmeratures(temps))
}
