package main

import "fmt"

func findTheWinner(n int, k int) int {
	players := make([]int, n)
	for i := 1; i <= n; i++ {
		players[i-1] = i
	}
	counter := 0
	i := 0
	for len(players) > 0 && len(players) > k {

		counter += k
		i = counter % (n)
		players = append(players[:i-1], players[i:]...)
		counter = i
		fmt.Println("array : ", players, i, counter)

		/*if counter == k {
			players = append(players[:i], players[i+1:]...)
			fmt.Println("array : ", players)
			counter = 1
		}
		if i == len(players)-1 {
			i = 0
		}
		i++
		counter++
		*/
	}
	return players[0]
}
func main() {

	fmt.Println("result : ", findTheWinner(5, 2))
}
