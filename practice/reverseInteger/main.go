package main

import "fmt"



func reverse(x int) int {
	ret := 0
	while x >0{
		ret = ret*10 + x %10
		x = x/10
	} 
	return ret
}


func main(){
	fmt.Println(reverse(123))
}