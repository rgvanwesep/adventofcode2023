package day1

import (
	"regexp"
)

func Reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FindValue(input string) int {
	if len(input) == 0 {
	return 0
	}

	var firstDigit, lastDigit int

	validDigit := regexp.MustCompile(`[1-9]|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`)
	validReverseDigit := regexp.MustCompile(`[1-9]|(eno)|(owt)|(eerht)|(ruof)|(evif)|(xis)|(neves)|(thgie)|(enin)`)

	digitMap := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	firstDigit = digitMap[validDigit.FindString(input)]
	lastDigit = digitMap[Reverse(validReverseDigit.FindString(Reverse(input)))]

	return firstDigit*10 + lastDigit
}

func Sum(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += FindValue(input)
	}
	return sum
}
