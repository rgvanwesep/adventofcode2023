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
		ranges = append(ranges, Range{r.Start, intersect.Start - r.Start})
	}
	if intersect.Start+intersect.Length < r.Start+r.Length {
		ranges = append(ranges, Range{intersect.Start + intersect.Length, r.Start + r.Length - intersect.Start - intersect.Length})
	}
	return ranges
}

type SeedRanges []Range

func NewSeedRanges(input string) SeedRanges {
	input = strings.Replace(input, seedsPrefix, "", 1)
	input = strings.Trim(input, " \n")
	ranges := make(SeedRanges, 0)
	split := strings.Split(input, " ")
	for i, j := 0, 1; j < len(split); i, j = i+2, j+2 {
		start, _ := strconv.Atoi(split[i])
		length, _ := strconv.Atoi(split[j])
		ranges = append(ranges, Range{start, length})
	}
	return ranges
}

type RangeMapItem struct {
	DstStart int
	SrcStart int
	Length   int
}

func NewRangeMapItem(input string) RangeMapItem {
	input = strings.TrimRight(input, "\n")
	split := strings.Split(input, " ")
	dstStart, _ := strconv.Atoi(split[0])
	srcStart, _ := strconv.Atoi(split[1])
	length, _ := strconv.Atoi(split[2])
	return RangeMapItem{dstStart, srcStart, length}
}

func (r RangeMapItem) Apply(input int) (int, bool) {
	if input < r.SrcStart || input >= r.SrcStart+r.Length {
		return input, false
	}
	return input - r.SrcStart + r.DstStart, true
}

func (r RangeMapItem) ApplyRange(input Range) []Range {
	intersection := Range{r.SrcStart, r.Length}.Intersect(input)
	difference := input.Subtract(intersection)
	ranges := make([]Range, 0)
	if intersection.Length > 0 {
		ranges = append(ranges, Range{intersection.Start - r.SrcStart + r.DstStart, intersection.Length})
	}
	return append(ranges, difference...)
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

func (r RangeMap) ApplyRange(input Range) []Range {
	outputs := make([]Range, 0)
	for _, item := range r.Items {
		outputs = append(outputs, item.ApplyRange(input)...)
	}
	if len(outputs) == 0 {
		return []Range{input}
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

type RangeAtlas struct {
	SeedRanges SeedRanges
	Maps       []RangeMap
}

func NewRangeAtlas(inputs []string) RangeAtlas {
	atlas := RangeAtlas{
		SeedRanges: NewSeedRanges(inputs[0]),
		Maps:       make([]RangeMap, 0),
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

func (a RangeAtlas) FindLocations() []Range {
	seedRanges := a.SeedRanges
	inputs := make([]Range, 0)
	outputs := make([]Range, 0)
	locations := make([]Range, 0)
	for _, seedRange := range seedRanges {
		inputs = append(inputs, seedRange)
		for _, rangeMap := range a.Maps {
			for _, input := range inputs {
				outputs = append(outputs, rangeMap.ApplyRange(input)...)
			}
			inputs = outputs
			outputs = make([]Range, 0)
		}
		minLocation := inputs[0]
		for _, input := range inputs[1:] {
			if input.Start < minLocation.Start {
				minLocation = input
			}
		}
		locations = append(locations, minLocation)
		inputs = make([]Range, 0)
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

func MinLocationFromRanges(inputs []string) int {
	atlas := NewRangeAtlas(inputs)
	locations := atlas.FindLocations()
	minLocation := locations[0]
	for _, location := range locations[1:] {
		if location.Start < minLocation.Start {
			minLocation = location
		}
	}
	return minLocation.Start
}
