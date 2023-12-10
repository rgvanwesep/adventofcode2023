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

type Node struct {
	Value      byte
	Coordinate Coordinate
	Neighbors  []*Node
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

type Coordinate struct {
	X int
	Y int
}

type Graph struct {
	Nodes       map[Coordinate]*Node
	Start       Coordinate
	Distances   map[Coordinate]int
	MaxDistance int
}

func NewGraph(lines []string) *Graph {
	g := &Graph{
		Nodes: make(map[Coordinate]*Node),
	}
	for y, line := range lines {
		line = strings.TrimRight(line, "\n")
		for x, char := range line {
			if char == ground {
				continue
			}
			coord := Coordinate{x, y}
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

func (g *Graph) Connect() {
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
			}
			if southOK && (south.Value == vertical || south.Value == northEast || south.Value == northWest) {
				node.Neighbors = append(node.Neighbors, south)
			}
			if eastOK && (east.Value == horizontal || east.Value == northWest || east.Value == southWest) {
				node.Neighbors = append(node.Neighbors, east)
			}
			if westOK && (west.Value == horizontal || west.Value == northEast || west.Value == southEast) {
				node.Neighbors = append(node.Neighbors, west)
			}
		}
	}
}

func (g *Graph) UpdateDistances() {
	var nextOk bool
	g.Distances = make(map[Coordinate]int)
	g.Distances[g.Start] = 0
	startNode := g.Nodes[g.Start]
	nodes := [2]*Node{startNode.Neighbors[0], startNode.Neighbors[1]}
	prevNodes := [2]*Node{startNode, startNode}
outer:
	for distance := 1; nodes[0] != startNode && nodes[1] != startNode; distance++ {
		for i, node := range nodes {
			coord := Coordinate{node.Coordinate.X, node.Coordinate.Y}
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
}

func FindFarthest(lines []string) int {
	g := NewGraph(lines)
	g.Connect()
	g.UpdateDistances()
	return g.MaxDistance
}
