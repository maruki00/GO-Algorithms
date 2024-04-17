package main

import (
	"fmt"
)

type Graph struct {
	Verteces []*Vertex
}

type Vertex struct {
	key       int
	adjancent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if g.contains(g.Verteces, k) {
		fmt.Println(fmt.Errorf("Vertex Alread Exists").Error())
		return
	}
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

func (g *Graph) contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func main() {

	g := &Graph{}
	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}
	g.Print()
}
