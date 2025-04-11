package main

import "fmt"

func divisorGame_(n int) bool {
	return n%2 == 0
}

func divisorGame(n int) bool {
	return n&1 == 0
}

func main() {
	n := 2
	fmt.Println("result : ", divisorGame(n))
}
