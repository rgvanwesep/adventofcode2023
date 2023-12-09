package day8

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	left  = 'L'
	right = 'R'
	start = "AAA"
	end   = "ZZZ"
)

var (
	directionRegex = regexp.MustCompile(`^(L|R)+$`)
	nodeRegex      = regexp.MustCompile(`^([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)$`)
)

type Directions string

func NewDirections(line string) (Directions, error) {
	line = strings.TrimRight(line, "\n")
	if !directionRegex.MatchString(line) {
		return "", fmt.Errorf("invalid directions")
	}
	return Directions(line), nil
}

type Node struct {
	Id    string
	Left  string
	Right string
}

func NewNode(line string) (Node, error) {
	line = strings.TrimRight(line, "\n")
	match := nodeRegex.FindStringSubmatch(line)
	if match == nil {
		return Node{}, fmt.Errorf("invalid node")
	}
	return Node{
		Id:    match[1],
		Left:  match[2],
		Right: match[3],
	}, nil
}

type Graph map[string]Node

func NewGraph(lines []string) Graph {
	graph := make(Graph)
	for _, line := range lines {
		node, err := NewNode(line)
		if err != nil {
			continue
		}
		graph[node.Id] = node
	}
	return graph
}

func (g Graph) Equals(other Graph) bool {
	if len(g) != len(other) {
		return false
	}
	for k, v := range g {
		if other[k] != v {
			return false
		}
	}
	return true
}

func CountSteps(lines []string) int {
	directions, err := NewDirections(lines[0])
	lenDirections := len(directions)
	if err != nil {
		return -1
	}
	graph := NewGraph(lines[1:])
	node := graph[start]
	index := 0
	steps := 0
	for {
		direction := directions[index]
		if direction == left {
			node = graph[node.Left]
		} else {
			node = graph[node.Right]
		}
		steps++
		if node.Id == end {
			break
		}
		index = (index + 1) % lenDirections
	}
	return steps
}
