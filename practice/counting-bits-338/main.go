pckage main

func countBits(num int) []int {
    bits := make([]int,num+1)
    for i:=1;i<num+1;i++ {
        bits[i] = bits[i&(i-1)]+1
    }   
    return bits
}

func main(){
	
}
