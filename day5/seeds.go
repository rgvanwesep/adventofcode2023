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

func (s SeedRanges) Seeds() Seeds {
	seeds := make(Seeds, 0)
	for _, r := range s {
		for i := 0; i < r.Length; i++ {
			seeds = append(seeds, r.Start+i)
		}
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

func (r RangeMapItem) ApplyRange(input Range) (Range, []Range) {
	intersection := Range{r.DstStart, r.Length}.Intersect(input)
	if intersection.Length == 0 {
		return Range{}, []Range{input}
	}
	mapped := Range{intersection.Start - r.SrcStart + r.DstStart, intersection.Length}
	unmapped := make([]Range, 0)
	for _, unmappedRange := range input.Subtract(intersection) {
		if unmappedRange.Length > 0 {
			unmapped = append(unmapped, unmappedRange)
		}
	}
	return mapped, unmapped
}

func (r RangeMapItem) ApplyRanges(inputs []Range) ([]Range, []Range) {
	mapped := make([]Range, 0)
	unmapped := make([]Range, 0)
	for _, input := range inputs {
		output, unmappedRanges := r.ApplyRange(input)
		if output.Length > 0 {
			mapped = append(mapped, output)
		}
		unmapped = append(unmapped, unmappedRanges...)
	}
	return mapped, unmapped
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

func (r RangeMap) ApplyRange(input Range) []Range {
	inputs := []Range{input}
	outputs := make([]Range, 0)
	for _, item := range r.Items {
		mapped, unmapped := item.ApplyRanges(inputs)
		outputs = append(outputs, mapped...)
		inputs = unmapped
	}
	return append(outputs, inputs...)
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

func NewAtlasFromRanges(inputs []string) Atlas {
	atlas := Atlas{
		Seeds: NewSeedRanges(inputs[0]).Seeds(),
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

func MinLocationFromRanges(inputs []string) int {
	atlas := NewAtlasFromRanges(inputs)
	locations := atlas.FindLocations()
	minLocation := locations[0]
	for _, location := range locations[1:] {
		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
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
	inputs := a.SeedRanges
	outputs := make([]Range, 0)
	for _, rangeMap := range a.Maps {
		for _, input := range inputs {
			outputs = append(outputs, rangeMap.ApplyRange(input)...)
		}
		inputs = outputs
		outputs = make([]Range, 0)
	}
	return inputs
}

func MinLocationFromRangeAtlas(inputs []string) int {
	atlas := NewRangeAtlas(inputs)
	locations := atlas.FindLocations()
	minLocation := locations[0].Start
	for _, location := range locations[1:] {
		if location.Start < minLocation {
			minLocation = location.Start
		}
	}
	return minLocation
}
