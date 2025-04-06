package main

func equalPairs(grid [][]int) int {
	visited := make(map[[2]int]bool)
	counter := 0
	n := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num1, ok1 := visited[[2]int{i, j}]
			if ok1 {
				continue
			}
			num2, ok2 := visited[[2]int{j, i}]
			if num1 == num2 && !ok2 {
				counter++
				visited[[2]int{i, j}] = true
			}
		}
	}
	return counter
}

func main() {

}
