package main

import (
	"fmt"
)

type Node struct {
	node      int
	adjancies []int
}

type Graph struct {
	adjancies map[byte][]byte
}

func (g *Graph) AddAdjancy(node, adj byte) {
	for _, nd := range g.adjancies[node] {
		if nd == adj {
			return
		}
	}
	g.adjancies[node] = append(g.adjancies[node], adj)
}

func DFS(start byte, g *Graph) {
	visited := make(map[byte]bool, len(g.adjancies))
	queue := []byte{start}
	fmt.Println(queue, visited)
}

func (g *Graph) Print() {
	for node, vals := range g.adjancies {
		fmt.Print("Node [", string(node), "] : ")
		for _, item := range vals {
			print(string(item), ", ")
		}
		fmt.Println("")
	}
}
func main() {
	g := &Graph{
		adjancies: make(map[byte][]byte),
	}
	g.AddAdjancy('a', 'b')
	g.AddAdjancy('b', 'b')
	g.AddAdjancy('c', 'b')
	g.AddAdjancy('a', 'c')
	g.Print()
}
