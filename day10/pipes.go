package day10

import (
	"log"
	"strings"
)

const (
	vertical   = '|'
	horizontal = '-'
	northEast  = 'L'
	northWest  = 'J'
	southWest  = '7'
	southEast  = 'F'
	ground     = '.'
	start      = 'S'
)

const (
	mainLoop = iota
	leftShoulder
	rightShoulder
)

type Node struct {
	Value      byte
	Coordinate Coordinate
	Neighbors  []*Node
	Class 	int
}

func (n *Node) Next(in *Node) (*Node, bool) {
	if len(n.Neighbors) != 2 {
		return nil, false
	}
	if n.Neighbors[0] == in {
		return n.Neighbors[1], true
	}
	return n.Neighbors[0], true
}

func (n *Node) Shoulders(in Coordinate) [2]Coordinate {
	var shoulders [2]Coordinate
	switch n.Value {
	case vertical:
		if in.Y > n.Coordinate.Y {
			shoulders[0] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y}
			shoulders[1] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y}
		} else {
			shoulders[0] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y}
			shoulders[1] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y}
		}
	case horizontal:
		if in.X < n.Coordinate.X {
			shoulders[0] = Coordinate{n.Coordinate.X, n.Coordinate.Y - 1}
			shoulders[1] = Coordinate{n.Coordinate.X, n.Coordinate.Y + 1}
		} else {
			shoulders[0] = Coordinate{n.Coordinate.X, n.Coordinate.Y + 1}
			shoulders[1] = Coordinate{n.Coordinate.X, n.Coordinate.Y - 1}
		}
	case northEast:
		if in.Y < n.Coordinate.Y {
			shoulders[0] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y - 1}
			shoulders[1] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y + 1}
		} else {
			shoulders[0] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y + 1}
			shoulders[1] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y - 1}
		}
	case northWest:
		if in.Y < n.Coordinate.Y {
			shoulders[0] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y + 1}
			shoulders[1] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y - 1}
		} else {
			shoulders[0] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y - 1}
			shoulders[1] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y + 1}
		}
	case southWest:
		if in.Y > n.Coordinate.Y {
			shoulders[0] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y + 1}
			shoulders[1] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y - 1}
		} else {
			shoulders[0] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y - 1}
			shoulders[1] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y + 1}
		}
	case southEast:
		if in.Y > n.Coordinate.Y {
			shoulders[0] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y - 1}
			shoulders[1] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y + 1}
		} else {
			shoulders[0] = Coordinate{n.Coordinate.X + 1, n.Coordinate.Y + 1}
			shoulders[1] = Coordinate{n.Coordinate.X - 1, n.Coordinate.Y - 1}
		}
	}
	return shoulders
}

type Coordinate struct {
	X int
	Y int
}

type Graph struct {
	Nodes       map[Coordinate]*Node
	Start       Coordinate
	Distances   map[Coordinate]int
	MainLoop    map[Coordinate]*Node
	Outsides	map[Coordinate]bool
	Shoulders   map[Coordinate][2]Coordinate
	MaxDistance int
}

func NewGraph(lines []string) *Graph {
	g := &Graph{
		Nodes: make(map[Coordinate]*Node),
	}
	for y, line := range lines {
		line = strings.TrimRight(line, "\n")
		for x, char := range line {
			coord := Coordinate{x, y}
			g.Outsides[coord] = true
			g.Nodes[coord] = &Node{
				Value:      byte(char),
				Coordinate: coord,
			}
			if char == start {
				g.Start = coord
			}
		}
	}
	return g
}

func (g *Graph) ConnectPipes() {
	var northAppended, southAppended, eastAppended, westAppended bool
	for coord, node := range g.Nodes {
		north, northOK := g.Nodes[Coordinate{coord.X, coord.Y - 1}]
		south, southOK := g.Nodes[Coordinate{coord.X, coord.Y + 1}]
		east, eastOK := g.Nodes[Coordinate{coord.X + 1, coord.Y}]
		west, westOK := g.Nodes[Coordinate{coord.X - 1, coord.Y}]
		switch node.Value {
		case vertical:
			node.Neighbors = []*Node{north, south}
		case horizontal:
			node.Neighbors = []*Node{east, west}
		case northEast:
			node.Neighbors = []*Node{north, east}
		case northWest:
			node.Neighbors = []*Node{north, west}
		case southWest:
			node.Neighbors = []*Node{south, west}
		case southEast:
			node.Neighbors = []*Node{south, east}
		case start:
			if northOK && (north.Value == vertical || north.Value == southEast || north.Value == southWest) {
				node.Neighbors = append(node.Neighbors, north)
				northAppended = true
			}
			if southOK && (south.Value == vertical || south.Value == northEast || south.Value == northWest) {
				node.Neighbors = append(node.Neighbors, south)
				southAppended = true
			}
			if eastOK && (east.Value == horizontal || east.Value == northWest || east.Value == southWest) {
				node.Neighbors = append(node.Neighbors, east)
				eastAppended = true
			}
			if westOK && (west.Value == horizontal || west.Value == northEast || west.Value == southEast) {
				node.Neighbors = append(node.Neighbors, west)
				westAppended = true
			}
			switch {
			case northAppended && southAppended:
				node.Value = vertical
			case eastAppended && westAppended:
				node.Value = horizontal
			case northAppended && eastAppended:
				node.Value = northEast
			case northAppended && westAppended:
				node.Value = northWest
			case southAppended && eastAppended:
				node.Value = southEast
			case southAppended && westAppended:
				node.Value = southWest
			}
		}
	}
}

func (g *Graph) MapMainLoop() {
	var nextOk bool
	g.Distances = make(map[Coordinate]int)
	g.Distances[g.Start] = 0
	startNode := g.Nodes[g.Start]
	g.MainLoop = map[Coordinate]*Node{g.Start: startNode}
	nodes := [2]*Node{startNode.Neighbors[0], startNode.Neighbors[1]}
	prevNodes := [2]*Node{startNode, startNode}
outer:
	for distance := 1; nodes[0] != startNode && nodes[1] != startNode; distance++ {
		for i, node := range nodes {
			coord := Coordinate{node.Coordinate.X, node.Coordinate.Y}
			g.MainLoop[coord] = node
			if _, ok := g.Distances[coord]; !ok {
				g.Distances[coord] = distance
			} else {
				g.MaxDistance = distance
				break outer
			}
			nodes[i], nextOk = node.Next(prevNodes[i])
			if !nextOk {
				log.Fatalf("Node %v has no next", node)
			}
			prevNodes[i] = node
		}
	}

	prev := startNode
	node := startNode.Neighbors[0]
	var next *Node
	for node != startNode {
		node.Class = mainLoop
		g.Outsides[node.Coordinate] = false
		g.Shoulders[node.Coordinate] = node.Shoulders(prev.Coordinate)
		next, nextOk = node.Next(prev)
		if !nextOk {
			log.Fatalf("Node %v has no next", node)
		}
		prev, node = node, next
	}
	g.Outsides[g.Start] = false
	g.Shoulders[g.Start] = startNode.Shoulders(prev.Coordinate)

}

func (g *Graph) ConnectOutsides() {
	for coord, outside := range g.Outsides {
		if !outside {
			continue
		}
		node := g.Nodes[coord]
		node.Neighbors = make([]*Node, 0)
		for _, x := range []int{-1, 0, 1} {
			for _, y := range []int{-1, 0, 1} {
				if x == 0 && y == 0 {
					continue
				}
				neighbor, ok := g.Nodes[Coordinate{coord.X + x, coord.Y + y}]
				if ok && g.Outsides[neighbor.Coordinate] {
					node.Neighbors = append(node.Neighbors, neighbor)
				}
			}
		}
	}
}

func FindFarthest(lines []string) int {
	g := NewGraph(lines)
	g.ConnectPipes()
	g.MapMainLoop()
	return g.MaxDistance
}
