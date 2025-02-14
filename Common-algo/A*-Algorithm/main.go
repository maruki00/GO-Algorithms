package main

type Node struct {
	X, Y int
	cost int
	prev *Node
}

func AStar(start, goal *Node, grid [][]*Node) []*Node {
	var closedSet []*Node
	var openSet = []*Node{start}
	start.cost = 0

	for len(openSet) > 0 {
		var current = openSet[0]
		if current == goal {
			var path []*Node
			for current != nil {
				path = append([]*Node{current}, path...)
				current = current.prev
			}
			return path
		}
		openSet = openSet[1:]
		closedSet = append(closedSet, current)
	}
	return nil
}

func mian() {

}
