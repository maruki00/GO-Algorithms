package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	occures := make(map[int]int)
	tracking := make(map[int]int)
	for _, num := range arr {
		occures[num]++
	}
	if len(occures) == 1 {
		return true
	}
	for _, v := range occures {
		tracking[v]++
	}
	for k, v := range tracking {
		fmt.Println(k, v)
		if v == k && v == 1 {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{26, 2, 16, 16, 5, 5, 26, 2, 5, 20, 20, 5, 2, 20, 2, 2, 20, 2, 16, 20, 16, 17, 16, 2, 16, 20, 26, 16}
	fmt.Println("result :", uniqueOccurrences(arr))
}
