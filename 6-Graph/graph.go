package main

type Graph struct {
	Verteces []*Vertex
}
type Vertex struct {
	key       int
	adjancent []*Vertex
}
