package day1

import "testing"

func TestFindValue(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}
	for _, c := range cases {
		result := FindValue(c.input)
		if result != c.expected {
			t.Errorf("FindValue(%q) == %d, expected %d", c.input, result, c.expected)
		}
	}
}

func TestSum(t *testing.T) {
	inputs := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	const expected = 142

	result := Sum(inputs)
	if result != expected {
		t.Errorf("Sum(%v) == %d, expected %d", inputs, result, expected)
	}
}
