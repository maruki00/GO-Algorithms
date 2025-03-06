package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	a := 534568765239652936
	b := 98
	fmt.Printf("The GCD of %d and %d is %d\n", a, b, gcd(a, b))
}

