package day7

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	ace   = 'A'
	king  = 'K'
	queen = 'Q'
	jack  = 'J'
	ten   = 'T'
	nine  = '9'
	eight = '8'
	seven = '7'
	six   = '6'
	five  = '5'
	four  = '4'
	three = '3'
	two   = '2'
)

const (
	fiveOfAKind  = 7
	fourOfAKind  = 6
	fullHouse    = 5
	threeOfAKind = 4
	twoPair      = 3
	onePair      = 2
	highCard     = 1
)

const handSize = 5

type Card byte

func (c Card) Rank() byte {
	switch c {
	case ace:
		return 14
	case king:
		return 13
	case queen:
		return 12
	case jack:
		return 11
	case ten:
		return 10
	case nine:
		return 9
	case eight:
		return 8
	case seven:
		return 7
	case six:
		return 6
	case five:
		return 5
	case four:
		return 4
	case three:
		return 3
	case two:
		return 2
	}
	return 0
}

func (c Card) LessThan(other Card) bool {
	return c.Rank() < other.Rank()
}

type Hand struct {
	Cards [handSize]Card
	Bid   int
}

func NewHand(line string) Hand {
	line = strings.TrimRight(line, "\n")
	split := strings.Split(line, " ")
	cards := [handSize]Card{}
	for i, card := range split[0] {
		cards[i] = Card(card)
	}
	bid, _ := strconv.Atoi(split[1])
	return Hand{Cards: cards, Bid: bid}
}

func (h Hand) Counts() map[Card]int {
	counts := make(map[Card]int)
	for _, card := range h.Cards {
		counts[card]++
	}
	return counts
}

func (h Hand) Type() int {
	counts := make([]int, 0)
	for _, count := range h.Counts() {
		counts = append(counts, count)
	}
	slices.Sort(counts)
	nUnique := len(counts)
	switch nUnique {
	case 1:
		return fiveOfAKind
	case 2:
		countPair := [2]int{counts[0], counts[1]}
		switch countPair {
		case [2]int{2, 3}:
			return fullHouse
		case [2]int{1, 4}:
			return fourOfAKind
		}
	case 3:
		countTriple := [3]int{counts[0], counts[1], counts[2]}
		switch countTriple {
		case [3]int{1, 1, 3}:
			return threeOfAKind
		case [3]int{1, 2, 2}:
			return twoPair
		}
	case 4:
		return onePair
	case 5:
		return highCard
	}
	return 0
}

func (h Hand) LessThan(other Hand) bool {
	if h.Type() < other.Type() {
		return true
	}
	if h.Type() > other.Type() {
		return false
	}
	for i := 0; i < handSize; i++ {
		if h.Cards[i] == other.Cards[i] {
			continue
		}
		return h.Cards[i].LessThan(other.Cards[i])
	}
	return false
}

type Hands []Hand

func NewHands(lines []string) Hands {
	hands := make(Hands, 0)
	for _, line := range lines {
		hands = append(hands, NewHand(line))
	}
	return hands
}

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Less(i, j int) bool {
	if h[i].Cards == h[j].Cards {
		return false
	}
	return h[i].LessThan(h[j])
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func Winnings(lines []string) int {
	winnings := 0
	hands := NewHands(lines)
	sort.Sort(hands)
	for i, hand := range hands {
		winnings += hand.Bid * (i + 1)
	}
	return winnings
}
