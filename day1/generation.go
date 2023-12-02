package day1

import "math/rand"

const (
	letters        = "qwertyuiopasdfghjklzxcvbnm"
	minDigitTokens = 1
)

var (
	digitTokens = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"twone",
		"eightwo",
		"oneight",
		"threeight",
		"fiveight",
		"nineight",
		"sevenine",
		"eighthree",
	}
	nLetters    = len(letters)
	letterWidth = BitWidth(nLetters)
	letterDepth = 63 / letterWidth
	letterMask  = 1<<letterWidth - 1
	nTokens     = len(digitTokens)
	tokenWidth  = BitWidth(nTokens)
	tokenDepth  = 63 / tokenWidth
	tokenMask   = 1<<tokenWidth - 1
)

func BitWidth(n int) int {
	if n == 0 {
		return 1
	}
	width := 0
	for n != 0 {
		n >>= 1
		width++
	}
	return width
}

type BitCache struct {
	source    *rand.Source
	bits      int64
	remaining int
	width     int
	depth     int
	mask      int
}

type Generator struct {
	MaxDigitTokens int
	MaxPadding     int
	Source         rand.Source
	tokenBits      BitCache
	letterBits     BitCache
	nTokenBits     BitCache
	nPaddingBits   BitCache
}

func NewGenerator(maxDigitTokens int, maxPadding int, source rand.Source) Generator {
	var (
		nTokenWidth   = BitWidth(maxDigitTokens)
		nTokenDepth   = 63 / nTokenWidth
		nTokenMask    = 1<<nTokenWidth - 1
		nPaddingWidth = BitWidth(maxPadding)
		nPaddingDepth = 63 / nPaddingWidth
		nPaddingMask  = 1<<nPaddingWidth - 1
	)
	return Generator{
		maxDigitTokens,
		maxPadding,
		source,
		BitCache{&source, source.Int63(), tokenDepth, tokenWidth, tokenDepth, tokenMask},
		BitCache{&source, source.Int63(), letterDepth, letterWidth, letterDepth, letterMask},
		BitCache{&source, source.Int63(), nTokenDepth, nTokenWidth, nTokenDepth, nTokenMask},
		BitCache{&source, source.Int63(), nPaddingDepth, nPaddingWidth, nPaddingDepth, nPaddingMask},
	}
}
