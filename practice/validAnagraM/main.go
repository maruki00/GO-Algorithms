package main

import "fmt"

func isAnagram(s string , t string ) bool {
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}

	for i:=0, j:=lt-1; i<ls && j>=0; i++, j--{
		if s[i] != t[j]{
			return false
		}
		return true
	}
}

func main(){
	fmt.Println("Result : ", isAnagram("anagram", "nagaram"))
}
