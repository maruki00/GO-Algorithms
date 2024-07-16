package main

import (
	"fmt"
	"hash/fnv"
)

var defaultCapacity uint64 = 1 << 10

type Node struct {
	key   any
	value any
	next  *Node
}

type HashMap struct {
	size     uint64
	capacity uint64
	table    []*Node
}

func New() *HashMap {
	return &HashMap{
		size:     defaultCapacity,
		capacity: defaultCapacity,
		table:    make([]*Node, defaultCapacity),
	}
}

func (hm *HashMap) getHash(key any) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

func (hm *HashMap) Add(key any, value any) {
	// hash := hm.getHash()
}

func main() {
	h := New()
	hashValue := h.getHash("helloworld")
	h.table[hashValue] = &Node{1, 1, nil}
	fmt.Println("hello nworld", hashValue)

}
