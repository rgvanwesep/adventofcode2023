package day7

import "testing"

func TestWinnings(t *testing.T) {
	cases := []struct {
		lines    []string
		expected int
	}{
		{
			[]string{
				"32T3K 765\n",
				"T55J5 684\n",
				"KK677 28\n",
				"KTJJT 220\n",
				"QQQJA 483",
			},
			6440,
		},
	}
	for _, c := range cases {
		winnings := Winnings(c.lines)
		if winnings != c.expected {
			t.Errorf("Winnings(%v) == %v, expected %v", c.lines, winnings, c.expected)
		}
	}
}
