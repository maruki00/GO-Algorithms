package main

import "strconv"

func fizzBuzz(n int) []string {
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			out[i-1] = "FizzBuzz"
			continue
		}
		if i%3 == 0 {
			out[i-1] = "Fizz"
			continue
		}
		if i%5 == 0 {
			out[i-1] = "Buzz"
			continue
		}
		out[i-1] = strconv.Itoa(i)
	}
	return out
}

func main() {

}
