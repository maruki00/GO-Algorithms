package main

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}

	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

func main() {

}
