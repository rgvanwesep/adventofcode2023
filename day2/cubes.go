package day2

import (
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type Set struct {
	Red   int
	Green int
	Blue  int
}

func NewSet(input string) Set {
	counts := strings.Split(input, ", ")
	countMap := make(map[string]int)
	for _, count := range counts {
		pair := strings.Split(count, " ")
		countMap[pair[1]], _ = strconv.Atoi(pair[0])
	}
	return Set{countMap["red"], countMap["green"], countMap["blue"]}
}

func (set Set) Check() bool {
	return set.Red <= maxRed && set.Green <= maxGreen && set.Blue <= maxBlue
}

type Game struct {
	Id   int
	Sets []Set
}

func NewGame(input string) Game {
	parts := strings.Split(input, ": ")

	gamePair := strings.Split(parts[0], " ")
	id, _ := strconv.Atoi(gamePair[1])

	sets := make([]Set, 0)
	for _, set := range strings.Split(parts[1], "; ") {
		sets = append(sets, NewSet(set))
	}
	return Game{id, sets}
}

func (game Game) Check() bool {
	for _, set := range game.Sets {
		if !set.Check() {
			return false
		}
	}
	return true
}

func Sum(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		if game := NewGame(strings.TrimRight(input, "\n")); game.Check() {
			sum += game.Id
		}
	}
	return sum
}
