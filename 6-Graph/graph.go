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
		err := fmt.Errorf("Vertex %v Alread Exists", k)
		fmt.Println(err.Error())
		return
	}
	g.Verteces = append(g.Verteces, &Vertex{key: k})
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if toVertex == nil || fromVertex == nil {
		err := fmt.Errorf("Invalid Edge from %v to %d Exists", from, to)
		fmt.Println(err.Error())
		return
	}
	fromVertex.adjancent = append(fromVertex.adjancent, toVertex)
}

func (g *Graph) Print() {
	for _, v := range g.Verteces {
		fmt.Printf("\n Vertex %v : ", v.key)
		for _, vk := range v.adjancent {
			fmt.Printf(" %v, ", vk.key)
		}
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.Verteces {
		if v.key == k {
			return g.Verteces[i]
		}
	}
	return nil
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
	g.AddEdge(3, 1)
	g.AddEdge(2, 1)
	g.Print()
}
