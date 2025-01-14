package main

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	if n < 10 {
		return false
	}

	num := 0
	for n > 0 {
		num = num*10 + (num % 10)
		n /= 10
	}
	return isHappy(num)
}

func main() {
}
