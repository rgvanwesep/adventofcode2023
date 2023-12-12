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

func TestCountInner(t *testing.T) {
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
			1,
		},
		{
			[]string{
				"..F7.\n",
				".FJ|.\n",
				"SJ.L7\n",
				"|F--J\n",
				"LJ...",
			},
			1,
		},
		{
			[]string{
				"F-...\n",
				"|S-7.\n",
				".|.|.\n",
				".L-J.\n",
				".....",
			},
			1,
		},
		{
			[]string{
				"...........\n",
				".S-------7.\n",
				".|F-----7|.\n",
				".||.....||.\n",
				".||.....||.\n",
				".|L-7.F-J|.\n",
				".|..|.|..|.\n",
				".L--J.L--J.\n",
				"...........",
			},
			4,
		},
		{
			[]string{
				"..........\n",
				".S------7.\n",
				".|F----7|.\n",
				".||OOOO||.\n",
				".||OOOO||.\n",
				".|L-7F-J|.\n",
				".|II||II|.\n",
				".L--JL--J.\n",
				"..........",
			},
			4,
		},
		{
			[]string{
				".F----7F7F7F7F-7....\n",
				".|F--7||||||||FJ....\n",
				".||.FJ||||||||L7....\n",
				"FJL7L7LJLJ||LJ.L-7..\n",
				"L--J.L7...LJS7F-7L7.\n",
				"....F-J..F7FJ|L7L7L7\n",
				"....L7.F7||L7|.L7L7|\n",
				".....|FJLJ|FJ|F7|.LJ\n",
				"....FJL-7.||.||||...\n",
				"....L---J.LJ.LJLJ...",
			},
			8,
		},
		{
			[]string{
				"FF7FSF7F7F7F7F7F---7\n",
				"L|LJ||||||||||||F--J\n",
				"FL-7LJLJ||||||LJL-77\n",
				"F--JF--7||LJLJ7F7FJ-\n",
				"L---JF-JLJ.||-FJLJJ7\n",
				"|F|F-JF---7F7-L7L|7|\n",
				"|FFJF7L7F-JF7|JL---7\n",
				"7-L-JL7||F7|L7F-7F7|\n",
				"L.L7LFJ|||||FJL7||LJ\n",
				"L7JLJL-JLJLJL--JLJ.L",
			},
			10,
		},
	}
	for _, c := range cases {
		d := CountInside(c.lines)
		if d != c.expected {
			t.Errorf("CountInside(%v) == %d, expected %d", c.lines, d, c.expected)
		}
	}
}
