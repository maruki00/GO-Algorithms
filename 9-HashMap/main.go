package main

import (
	"fmt"
	"hash/fnv"
)

/*
var defaultCapacity uint64 = 10

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

func (hm *HashMap) getNodeByHash(key uint64) *Node {
	return hm.table[key]
}
func (hm *HashMap) Put(key any, value any) {
	hash := hm.getHash(key)
	node := hm.getNodeByHash(hash)
	fmt.Println(node.table)
	return
	table := node
	for table.next != nil {
		table = table.next
	}
	table.next = &Node{key, value, nil}
}
*/

var defaultCapacity uint64 = 1 << 10

type node struct {
	key   any
	value any
	next  *node
}

// HashMap is golang implementation of hashmap
type HashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

// New return new HashMap instance
func New() *HashMap {
	return &HashMap{
		capacity: defaultCapacity,
		table:    make([]*node, defaultCapacity),
	}
}

// Make creates a new HashMap instance with input size and capacity
func Make(size, capacity uint64) HashMap {
	return HashMap{
		size:     size,
		capacity: capacity,
		table:    make([]*node, capacity),
	}
}

// Get returns value associated with given key
func (hm *HashMap) Get(key any) any {
	node := hm.getNodeByHash(hm.hash(key))

	if node != nil {
		return node.value
	}

	return nil
}

// Put puts new key value in hashmap
func (hm *HashMap) Put(key any, value any) any {
	return hm.putValue(hm.hash(key), key, value)
}

// Contains checks if given key is stored in hashmap
func (hm *HashMap) Contains(key any) bool {
	node := hm.getNodeByHash(hm.hash(key))
	return node != nil
}

func (hm *HashMap) putValue(hash uint64, key any, value any) any {
	if hm.capacity == 0 {
		hm.capacity = defaultCapacity
		hm.table = make([]*node, defaultCapacity)
	}

	node := hm.getNodeByHash(hash)

	if node == nil {
		hm.table[hash] = newNode(key, value)

	} else if node.key == key {
		hm.table[hash] = newNodeWithNext(key, value, node)
		return value

	} else {
		hm.resize()
		return hm.putValue(hash, key, value)
	}

	hm.size++

	return value

}

func (hm *HashMap) getNodeByHash(hash uint64) *node {
	return hm.table[hash]
}

func (hm *HashMap) resize() {
	hm.capacity <<= 1

	tempTable := hm.table

	hm.table = make([]*node, hm.capacity)

	for i := 0; i < len(tempTable); i++ {
		node := tempTable[i]
		if node == nil {
			continue
		}

		hm.table[hm.hash(node.key)] = node
	}
}

func newNode(key any, value any) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func newNodeWithNext(key any, value any, next *node) *node {
	return &node{
		key:   key,
		value: value,
		next:  next,
	}
}

func (hm *HashMap) hash(key any) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

func main() {
	h := New()
	h.Put(1, 1234)
	h.Put(1234, 234)
	//hashValue := h.getHash("helloworld")
	//h.table[hashValue] = &Node{1, 1, nil}
	fmt.Println("hello nworld", h.Contains(1))

}
