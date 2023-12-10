package day5

import (
	"log"
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

func (rs RangeSet) Bracket(r Range) (start, end int) {
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
	return
}

func (rs RangeSet) CoveredLeft(r Range) bool {
	if r.Length() == 0 {
		return false
	}
	if len(rs) == 0 {
		log.Fatal("RangeSet is empty")
	}
	return r.Start <= rs[0].Start
}

func (rs RangeSet) CoveredRight(r Range) bool {
	if r.Length() == 0 {
		return false
	}
	if len(rs) == 0 {
		log.Fatal("RangeSet is empty")
	}
	return r.End >= rs[len(rs)-1].End
}

func (rs RangeSet) Covered(r Range) bool {
	if r.Length() == 0 {
		return false
	}
	if len(rs) == 0 {
		log.Fatal("RangeSet is empty")
	}
	return rs.CoveredLeft(r) && rs.CoveredRight(r)
}

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

	start, end := rs.Bracket(r)
	result := make(RangeSet, 0)
	result = append(result, rs[:start]...)
	if start > end {
		result = append(result, r)
	} else if rs[start : end+1].Covered(r) {
		result = append(result, Range{r.Start, r.End})
	} else if rs[start : end+1].CoveredLeft(r) {
		result = append(result, Range{r.Start, rs[end].End})
	} else if rs[start : end+1].CoveredRight(r) {
		result = append(result, Range{rs[start].Start, r.End})
	} else {
		result = append(result, Range{rs[start].Start, rs[end].End})
	}
	result = append(result, rs[end+1:]...)
	return result
}

func (rs RangeSet) Union(other RangeSet) RangeSet {
	if len(rs) == 0 {
		return other
	}
	if len(other) == 0 {
		return rs
	}
	union := make(RangeSet, 0)
	for _, r := range rs {
		union = union.Add(r)
	}
	for _, r := range other {
		union = union.Add(r)
	}
	return union
}

func (rs RangeSet) Slice(start, end int) RangeSet {
	return rs[start:end]
}

func (rs RangeSet) Intersect(other RangeSet) RangeSet {
	if len(rs) == 0 || len(other) == 0 {
		return RangeSet{}
	}
	if rs[0].Start > other[len(other)-1].End || rs[len(rs)-1].End < other[0].Start {
		return RangeSet{}
	}
	intersection := make(RangeSet, 0)
	for _, r := range other {
		start, end := rs.Bracket(r)
		if start > end {
			continue
		} else if rs[start : end+1].Covered(r) {
			intersection = append(intersection, rs[start:end+1]...)
		} else if rs[start : end+1].CoveredLeft(r) {
			intersection = append(intersection, rs[start:end]...)
			intersection = append(intersection, Range{rs[end].Start, r.End})
		} else if rs[start : end+1].CoveredRight(r) {
			intersection = append(intersection, Range{r.Start, rs[start].End})
			intersection = append(intersection, rs[start+1:end+1]...)
		} else {
			intersection = append(intersection, Range{r.Start, rs[start].End})
			if start+1 < end {
				intersection = append(intersection, rs[start+1:end]...)
			}
			intersection = append(intersection, Range{rs[end].Start, r.End})
		}
		rs = rs[end:]
		if len(rs) == 0 {
			break
		}
	}
	return intersection
}

func (rs RangeSet) Subtract(other RangeSet) RangeSet {
	if len(rs.Intersect(other)) == 0 {
		return rs
	}
	subtraction := make(RangeSet, 0)
	for _, r := range other {
		start, end := rs.Bracket(r)
		subtraction = append(subtraction, rs[:start]...)
		if rs[start : end+1].Covered(r) {
			continue
		} else if rs[start : end+1].CoveredLeft(r) {
			newR := Range{r.End, rs[end].End}
			if newR.Length() > 0 {
				subtraction = append(subtraction, newR)
			}
		} else if rs[start : end+1].CoveredRight(r) {
			newR := Range{rs[start].Start, r.Start}
			if newR.Length() > 0 {
				subtraction = append(subtraction, Range{rs[start].Start, r.Start})
			}
		} else {
			newR := Range{rs[start].Start, r.Start}
			if newR.Length() > 0 {
				subtraction = append(subtraction, Range{rs[start].Start, r.Start})
			}
			newR = Range{r.End, rs[end].End}
			if newR.Length() > 0 {
				subtraction = append(subtraction, Range{r.End, rs[end].End})
			}
		}
		if len(rs) == 0 {
			break
		}
		rs = rs[end+1:]
		if len(rs) == 0 {
			break
		}
	}
	if len(rs) > 0 {
		subtraction = append(subtraction, rs[1:]...)
	}
	return subtraction
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

func (r RangeMapItem) ApplySet(input RangeSet) RangeSet {
	shift := r.DstStart - r.SrcStart
	srcSet := RangeSet{r.SrcRange()}
	intersection := srcSet.Intersect(input)
	if len(intersection) == 0 {
		return input
	}
	subtraction := srcSet.Subtract(input)
	output := make(RangeSet, 0)
	for _, r := range intersection {
		output = output.Add(Range{r.Start + shift, r.End + shift})
	}
	for _, r := range subtraction {
		output = output.Add(r)
	}
	return output
}

func NewRangeMapItem(input string) RangeMapItem {
	input = strings.TrimRight(input, "\n")
	split := strings.Split(input, " ")
	dstStart, _ := strconv.Atoi(split[0])
	srcStart, _ := strconv.Atoi(split[1])
	length, _ := strconv.Atoi(split[2])
	return RangeMapItem{dstStart, srcStart, length}
}

func (r RangeMapItem) SrcRange() Range {
	return Range{r.SrcStart, r.SrcStart + r.Length}
}

func (r RangeMapItem) DstRange() Range {
	return Range{r.DstStart, r.DstStart + r.Length}
}

type RangeMap []RangeMapItem

func (r RangeMap) SrcSet() RangeSet {
	rs := make(RangeSet, 0)
	for _, item := range r {
		rs = rs.Add(item.SrcRange())
	}
	return rs
}

func (r RangeMap) DstSet() RangeSet {
	rs := make(RangeSet, 0)
	for _, item := range r {
		rs = rs.Add(item.DstRange())
	}
	return rs
}

func (r RangeMap) Apply(input RangeSet) RangeSet {
	intersection := r.SrcSet().Intersect(input)
	output := make(RangeSet, 0)
	for _, item := range r {
		for _, r := range item.ApplySet(intersection) {
			output = output.Add(r)
		}
	}
	for _, r := range input.Subtract(r.SrcSet()) {
		output = output.Add(r)
	}
	return output
}

type NamedRangeMap struct {
	Header string
	Items  RangeMap
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
	Seeds RangeSet
	Maps  []NamedRangeMap
}

func NewAtlas(inputs []string) Atlas {
	seeds := NewSeeds(inputs[0])
	rangeSet := make(RangeSet, 0)
	for _, seed := range seeds {
		rangeSet = rangeSet.Add(Range{seed, seed + 1})
	}
	atlas := Atlas{
		Seeds: rangeSet,
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
	locations := make([]int, 0)
	inputs := a.Seeds
	for _, rangeMap := range a.Maps {
		outputs := rangeMap.Items.Apply(inputs)
		inputs = outputs
	}
	for _, input := range inputs {
		minLocation := input.Start
		locations = append(locations, minLocation)
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
