package day5

import (
	"reflect"
	"testing"
)

func TestNewAtlas(t *testing.T) {
	cases := []struct {
		inputs []string
		atlas  Atlas
	}{
		{
			inputs: []string{
				"seeds: 79 14 55 13\n",
				"\n",
				"seed-to-soil map:\n",
				"50 98 2\n",
				"52 50 48\n",
				"\n",
				"soil-to-fertilizer map:\n",
				"0 15 37\n",
				"37 52 2\n",
				"39 0 15\n",
				"\n",
				"fertilizer-to-water map:\n",
				"49 53 8\n",
				"0 11 42\n",
				"42 0 7\n",
				"57 7 4\n",
				"\n",
				"water-to-light map:\n",
				"88 18 7\n",
				"18 25 70\n",
				"\n",
				"light-to-temperature map:\n",
				"45 77 23\n",
				"81 45 19\n",
				"68 64 13\n",
				"\n",
				"temperature-to-humidity map:\n",
				"0 69 1\n",
				"1 0 69\n",
				"\n",
				"humidity-to-location map:\n",
				"60 56 37\n",
				"56 93 4",
			},
			atlas: Atlas{
				Seeds: Seeds{79, 14, 55, 13},
				Maps: []RangeMap{
					{
						Header: "seed-to-soil map:",
						Items: []RangeMapItem{
							{50, 98, 2},
							{52, 50, 48},
						},
					},
					{
						Header: "soil-to-fertilizer map:",
						Items: []RangeMapItem{
							{0, 15, 37},
							{37, 52, 2},
							{39, 0, 15},
						},
					},
					{
						Header: "fertilizer-to-water map:",
						Items: []RangeMapItem{
							{49, 53, 8},
							{0, 11, 42},
							{42, 0, 7},
							{57, 7, 4},
						},
					},
					{
						Header: "water-to-light map:",
						Items: []RangeMapItem{
							{88, 18, 7},
							{18, 25, 70},
						},
					},
					{
						Header: "light-to-temperature map:",
						Items: []RangeMapItem{
							{45, 77, 23},
							{81, 45, 19},
							{68, 64, 13},
						},
					},
					{
						Header: "temperature-to-humidity map:",
						Items: []RangeMapItem{
							{0, 69, 1},
							{1, 0, 69},
						},
					},
					{
						Header: "humidity-to-location map:",
						Items: []RangeMapItem{
							{60, 56, 37},
							{56, 93, 4},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		atlas := NewAtlas(c.inputs)
		if !reflect.DeepEqual(atlas, c.atlas) {
			t.Errorf("NewAtlas(%v) == %v, want %v", c.inputs, atlas, c.atlas)
		}
	}
}
