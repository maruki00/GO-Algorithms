package main 


func numToArray(num int ) int[]{
    var arr int[]
    for num > 0 {
		arr = append(arr, num%10)
		num = (int)(num/10)
    }
	return arr
}

func isPalindrome(x int) bool {
    if x <= 10{
        return false
    }


}

func main(){

}

