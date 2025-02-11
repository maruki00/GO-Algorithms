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
	for len(queue) > 0 {
		item := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if _, ok := visited[item]; ok {
			continue
		}
		visited[item] = true
		fmt.Println("Node visited : ", string(item))
		queue = append(queue, g.adjancies[item]...)

	}
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
	g.AddAdjancy('1', '2')
	g.AddAdjancy('1', '4')
	g.AddAdjancy('2', '3')
	g.AddAdjancy('4', '3')
	g.Print()
	DFS('a', g)
}
