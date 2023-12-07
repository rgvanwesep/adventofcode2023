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

type Range struct {
	Start  int
	Length int
}

func (r Range) Intersect(other Range) Range {
	start := r.Start
	if other.Start > start {
		start = other.Start
	}
	end := r.Start + r.Length
	if other.Start+other.Length < end {
		end = other.Start + other.Length
	}
	if end < start {
		return Range{start, 0}
	}
	return Range{start, end - start}
}

func (r Range) Subtract(other Range) []Range {
	intersect := r.Intersect(other)
	if intersect.Length == 0 {
		return []Range{r}
	}
	ranges := make([]Range, 0)
	if intersect.Start > r.Start {
		rtmp := Range{r.Start, intersect.Start - r.Start}
		if rtmp.Length > 0 {
			ranges = append(ranges, rtmp)
		}
	}
	if intersect.Start+intersect.Length < r.Start+r.Length {
		rtmp := Range{intersect.Start + intersect.Length, r.Start + r.Length - intersect.Start - intersect.Length}
		if rtmp.Length > 0 {
			ranges = append(ranges, rtmp)
		}
	}
	return ranges
}

func (r Range) Union(other Range) []Range {
	if r.Length == 0 {
		return []Range{other}
	}
	if other.Length == 0 {
		return []Range{r}
	}
	intersect := r.Intersect(other)
	if intersect.Length == 0 {
		return []Range{r, other}
	}
	start := r.Start
	if other.Start < start {
		start = other.Start
	}
	end := r.Start + r.Length
	if other.Start+other.Length > end {
		end = other.Start + other.Length
	}
	return []Range{{start, end - start}}
}

type RangeSet []Range

func NewRangeSet(ranges []Range) RangeSet {
	rangeSet := make(RangeSet, 0)
	for i := 0; i < len(ranges); i++ {
		for j := i + 1; j < len(ranges); j++ {
			ranges = append(ranges, ranges[i].Union(ranges[j])...)
		}
	}
	return rangeSet
}

func (r RangeSet) Intersect(other RangeSet) RangeSet {
	ranges := make(RangeSet, 0)
	for _, r1 := range r {
		for _, r2 := range other {
			intersection := r1.Intersect(r2)
			if intersection.Length > 0 {
				ranges = append(ranges, intersection)
			}
		}
	}
	return ranges
}

func (r RangeSet) Subtract(other RangeSet) RangeSet {
	ranges := make(RangeSet, 0)
	for _, r1 := range r {
		for _, r2 := range other {
			ranges = append(ranges, r1.Subtract(r2)...)
		}
	}
	return ranges
}

type Seeds []int

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

type RangeMap []RangeMapItem

func (r RangeMap) Apply(RangeSet) RangeSet {
	ranges := make(RangeSet, 0)
	return ranges
}

func NewRangeMapItem(input string) RangeMapItem {
	input = strings.TrimRight(input, "\n")
	split := strings.Split(input, " ")
	dstStart, _ := strconv.Atoi(split[0])
	srcStart, _ := strconv.Atoi(split[1])
	length, _ := strconv.Atoi(split[2])
	return RangeMapItem{dstStart, srcStart, length}
}

type NamedRangeMap struct {
	Header string
	Items  []RangeMapItem
}

func NewNamedRangeMap(inputs []string) NamedRangeMap {
	header := strings.TrimRight(inputs[0], "\n")
	items := make([]RangeMapItem, 0)
	for _, input := range inputs[1:] {
		if strings.TrimRight(input, "\n") == "" {
			break
		}
		items = append(items, NewRangeMapItem(input))
	}
	return NamedRangeMap{header, items}
}

func (r NamedRangeMap) Apply(input int) []int {
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
	Maps  []NamedRangeMap
}

func NewAtlas(inputs []string) Atlas {
	atlas := Atlas{
		Seeds: NewSeeds(inputs[0]),
		Maps:  make([]NamedRangeMap, 0),
	}
	for i := 1; i < len(inputs); {
		if strings.HasPrefix(inputs[i], seedToSoilHeader) {
			rangeMap := NewNamedRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], soilToFertilizerHeader) {
			rangeMap := NewNamedRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], fertilizerToWaterHeader) {
			rangeMap := NewNamedRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], waterToLightHeader) {
			rangeMap := NewNamedRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], lightToTemperatureHeader) {
			rangeMap := NewNamedRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], temperatureToHumidityHeader) {
			rangeMap := NewNamedRangeMap(inputs[i:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap.Items)
		} else if strings.HasPrefix(inputs[i], humidityToLocationHeader) {
			rangeMap := NewNamedRangeMap(inputs[i:])
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
