package day4

import (
	"strconv"
	"strings"
)

type Set map[int]bool

func (s *Set) Add(i int) {
	(*s)[i] = true
}

func (s *Set) Intersect(other Set) Set {
	result := make(Set)
	for k := range *s {
		if other[k] {
			result.Add(k)
		}
	}
	return result
}

type Stack []int

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	i := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return i, true
}

type Card struct {
	Id      int
	Winners Set
	Picks   Set
}

func NewCard(input string) Card {
	card := Card{
		Id:      0,
		Winners: make(Set),
		Picks:   make(Set),
	}
	colonSplit := strings.Split(input, ":")
	spaceSplit := strings.Split(colonSplit[0], " ")
	card.Id, _ = strconv.Atoi(spaceSplit[len(spaceSplit)-1])
	pipeSplit := strings.Split(strings.Trim(colonSplit[1], " "), "|")
	winnerString := strings.Trim(pipeSplit[0], " ")
	pickString := strings.Trim(pipeSplit[1], " ")
	for _, winner := range strings.Split(winnerString, " ") {
		if winner != "" {
			winnerInt, _ := strconv.Atoi(winner)
			card.Winners.Add(winnerInt)
		}
	}
	for _, pick := range strings.Split(pickString, " ") {
		if pick != "" {
			pickInt, _ := strconv.Atoi(pick)
			card.Picks.Add(pickInt)
		}
	}
	return card
}

func (c Card) Score() int {
	nMatches := c.NumMatches()
	if nMatches == 0 {
		return 0
	}
	return 1 << (nMatches - 1)
}

func (c Card) NumMatches() int {
	return len(c.Winners.Intersect(c.Picks))
}

type Cards struct {
	CardMap map[int]Card
	IdStack Stack
	Count   int
}

func NewCards() Cards {
	return Cards{
		CardMap: make(map[int]Card),
		IdStack: Stack{},
	}
}

func (c *Cards) Add(card Card) {
	c.CardMap[card.Id] = card
	c.IdStack.Push(card.Id)
	c.Count++
}

func (c *Cards) Pop() (Card, bool) {
	id, ok := c.IdStack.Pop()
	if !ok {
		return Card{}, false
	}
	card := c.CardMap[id]
	for i := 0; i < card.NumMatches(); i++ {
		c.IdStack.Push(id + i + 1)
		c.Count++
	}
	return card, true
}

func (c *Cards) PopAll() []Card {
	cards := make([]Card, 0)
	for {
		card, ok := c.Pop()
		if !ok {
			break
		}
		cards = append(cards, card)
	}
	return cards
}

func Sum(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		input = strings.TrimRight(input, "\n")
		sum += NewCard(input).Score()
	}
	return sum
}

func SumCards(inputs []string) int {
	cards := NewCards()
	for _, input := range inputs {
		input = strings.TrimRight(input, "\n")
		cards.Add(NewCard(input))
	}
	cards.PopAll()
	return cards.Count
}
