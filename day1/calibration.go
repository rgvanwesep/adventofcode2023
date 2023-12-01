package day1

import (
	"regexp"
	"strings"
)

func FindValue(input string) int {
	if len(input) == 0 {
		return 0
	}

	var firstDigit, lastDigit int

	validDigit := regexp.MustCompile(`[1-9]|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`)

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

	specialCases := map[string]string{
		"twone":   "21",
		"eightwo": "82",
		"oneight": "18",
		"threeight": "38",
		"fiveight": "58",
		"nineight": "98",
	}

	for old, new := range specialCases {
		input = strings.ReplaceAll(input, old, new)
	}
	digits := validDigit.FindAllString(input, -1)
	firstDigit = digitMap[digits[0]]
	lastDigit = digitMap[digits[len(digits)-1]]

	return firstDigit*10 + lastDigit
}

func Sum(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += FindValue(input)
	}
	return sum
}
