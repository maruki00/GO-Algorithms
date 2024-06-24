package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := x
	reversed := 0
	for num > 0 {
		tmp := num % 10
		reversed = reversed*10 + tmp
		num = num / 10
	}
	return reversed == x
}

func main() {

}
