package main

type NodeDouble struct {
	previous *NodeDouble
	next     *NodeDouble
	item     any
}

type NodeSimple struct {
	next *NodeSimple
	item any
}
