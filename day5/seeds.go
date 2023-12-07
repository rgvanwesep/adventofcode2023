package day5

import (
	"strconv"
	"strings"
)

const (
	seedsPrefix                 = "seeds:"
	seedToSoilHeader            = "seed-to-soil map:"
	soilToFertilizerHeader      = "soil-to-fertilizer map:"
	fertilizerToWaterHeader     = "fertilizer-to-water map:"
	waterToLightHeader          = "water-to-light map:"
	lightToTemperatureHeader    = "light-to-temperature map:"
	temperatureToHumidityHeader = "temperature-to-humidity map:"
	humidityToLocationHeader    = "humidity-to-location map:"
)

type Seeds []int

type Range struct {
	Start int
	End   int
}

func (r Range) Length() int {
	return r.End - r.Start
}

type RangeSet []Range

func (rs RangeSet) Add(r Range) RangeSet {
	if len(rs) == 0 {
		if r.Length() == 0 {
			return rs
		}
		return RangeSet{r}
	}
	if r.Length() == 0 {
		return rs
	}
	if rs[0].Start > r.End {
		return append(RangeSet{r}, rs...)
	}
	if rs[len(rs)-1].End < r.Start {
		return append(rs, r)
	}
	var start, end int
	for i := 0; i < len(rs); i++ {
		if rs[i].End >= r.Start {
			start = i
			break
		}
	}
	for i := len(rs) - 1; i >= 0; i-- {
		if rs[i].Start <= r.End {
			end = i
			break
		}
	}
	result := make(RangeSet, 0)
	result = append(result, rs[:start]...)
	if start == end {
		if rs[start].Start < r.Start {
			if r.End < rs[start].End {
				result = append(result, Range{rs[start].Start, rs[start].End})
			} else {
				result = append(result, Range{rs[start].Start, r.End})
			}
		} else {
			if r.End < rs[start].End {
				result = append(result, Range{r.Start, rs[start].End})
			} else {
				result = append(result, Range{r.Start, r.End})
			}
		}
	} else {
		if rs[start].Start < r.Start {
			if r.End < rs[end].End {
				result = append(result, Range{rs[start].Start, rs[end].End})
			} else {
				result = append(result, Range{rs[start].Start, r.End})
			}
		} else {
			if r.End < rs[end].End {
				result = append(result, Range{r.Start, rs[end].End})
			} else {
				result = append(result, Range{r.Start, r.End})
			}
		}
	}
	result = append(result, rs[end+1:]...)
	return result
}

func NewRangeSet(ranges []Range) RangeSet {
	rs := make(RangeSet, 0)
	for _, r := range ranges {
		rs = rs.Add(r)
	}
	return rs
}

func NewSeeds(input string) Seeds {
	input = strings.Replace(input, seedsPrefix, "", 1)
	input = strings.Trim(input, " \n")
	seeds := make(Seeds, 0)
	for _, seed := range strings.Split(input, " ") {
		i, _ := strconv.Atoi(seed)
		seeds = append(seeds, i)
	}
	return seeds
}

type RangeMapItem struct {
	DstStart int
	SrcStart int
	Length   int
}

func (r RangeMapItem) Apply(input int) (int, bool) {
	if input < r.SrcStart || input >= r.SrcStart+r.Length {
		return input, false
	}
	return input - r.SrcStart + r.DstStart, true
}

func NewRangeMapItem(input string) RangeMapItem {
	input = strings.TrimRight(input, "\n")
	split := strings.Split(input, " ")
	dstStart, _ := strconv.Atoi(split[0])
	srcStart, _ := strconv.Atoi(split[1])
	length, _ := strconv.Atoi(split[2])
	return RangeMapItem{dstStart, srcStart, length}
}

type RangeMap struct {
	Header string
	Items  []RangeMapItem
}

func NewRangeMap(inputs []string) RangeMap {
	header := strings.TrimRight(inputs[0], "\n")
	items := make([]RangeMapItem, 0)
	for _, input := range inputs[1:] {
		if strings.TrimRight(input, "\n") == "" {
			break
		}
		items = append(items, NewRangeMapItem(input))
	}
	return RangeMap{header, items}
}

func (r RangeMap) Apply(input int) []int {
	outputs := make([]int, 0)
	for _, item := range r.Items {
		if output, ok := item.Apply(input); ok {
			outputs = append(outputs, output)
		}
	}
	if len(outputs) == 0 {
		return []int{input}
	}
	return outputs
}

type Atlas struct {
	Seeds Seeds
	Maps  []RangeMap
}

func NewAtlas(inputs []string) Atlas {
	atlas := Atlas{
		Seeds: NewSeeds(inputs[0]),
		Maps:  make([]RangeMap, 0),
	}
	for i := 1; i < len(inputs); {
		if strings.HasPrefix(inputs[i], seedToSoilHeader) {
			rangeMap := NewRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], soilToFertilizerHeader) {
			rangeMap := NewRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], fertilizerToWaterHeader) {
			rangeMap := NewRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], waterToLightHeader) {
			rangeMap := NewRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], lightToTemperatureHeader) {
			rangeMap := NewRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], temperatureToHumidityHeader) {
			rangeMap := NewRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], humidityToLocationHeader) {
			rangeMap := NewRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else {
			i++
		}
	}
	return atlas
}

func (a Atlas) FindLocations() []int {
	seeds := a.Seeds
	inputs := make([]int, 0)
	outputs := make([]int, 0)
	locations := make([]int, 0)
	for _, seed := range seeds {
		inputs = append(inputs, seed)
		for _, rangeMap := range a.Maps {
			for _, input := range inputs {
				outputs = append(outputs, rangeMap.Apply(input)...)
			}
			inputs = outputs
			outputs = make([]int, 0)
		}
		minLocation := inputs[0]
		for _, input := range inputs[1:] {
			if input < minLocation {
				minLocation = input
			}
		}
		locations = append(locations, minLocation)
		inputs = make([]int, 0)
	}
	return locations
}

func MinLocation(inputs []string) int {
	atlas := NewAtlas(inputs)
	locations := atlas.FindLocations()
	minLocation := locations[0]
	for _, location := range locations[1:] {
		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}
