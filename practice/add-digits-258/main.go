pckage main



func addDigits(num int) int {
	result := 0
	var solution := func(n int)
	solution = func(n int) int {
		result := 0
		if n <=9 {
			return n
		}
		for n>0 {
			result = (result*10) + int(n%10) 
			n/= 10
		}
		return solution(result)
	}
}

func main(){

}
