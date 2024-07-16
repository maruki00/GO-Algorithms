package main

import (
	"fmt"
	"hash/fnv"
)

type Node struct {
	key any
	value any
	next *Node
}

type HashMap strut {
	size int
	capacity int
	table *[]Node
}

func New() *HashMap{
	return &HashMap{
		size: 0,
		capacity: 0,
		table: make([]*Node),
	}
}




func (obj *HashMap) getHash(key any) int64 {
	r := fnv.New64a()
	_, _ = r.Write([]byte(fmt.Printf("%v", key)))

	hashValue := r.Sum64()
	return (obj.capacity - 1) & (hashValue ^ (hashValue >> 16))

}


func main(){

}
