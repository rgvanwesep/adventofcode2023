package day10

import "testing"

func TestFindFarthest(t *testing.T) {
	cases := []struct {
		lines    []string
		expected int
	}{
		{
			[]string{
				".....\n",
				".S-7.\n",
				".|.|.\n",
				".L-J.\n",
				".....",
			},
			4,
		},
		{
			[]string{
				"..F7.\n",
				".FJ|.\n",
				"SJ.L7\n",
				"|F--J\n",
				"LJ...",
			},
			8,
		},
		{
			[]string{
				"F-...\n",
				"|S-7.\n",
				".|.|.\n",
				".L-J.\n",
				".....",
			},
			4,
		},
	}
	for _, c := range cases {
		d := FindFarthest(c.lines)
		if d != c.expected {
			t.Errorf("FindFarthest(%v) == %d, expected %d", c.lines, d, c.expected)
		}
	}
}
