package main

import "fmt"

type Graph struct {
	Verteces []*Vertex
}
type Vertex struct {
	key       int
	adjancent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	g.Verteces = append(g.Verteces, &Vertex{key: k})
}

func (g *Graph) Print() {
	for _, v := range g.Verteces {
		fmt.Printf("\n Vertex %v : ", v.key)
		for _, vk := range v.adjancent {
			fmt.Printf(" %v, ", vk.key)
		}
	}
}
