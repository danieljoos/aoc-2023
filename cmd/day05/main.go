package main

import (
	"math"
	"regexp"
	"slices"
	"strings"
	"sync"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
	"github.com/schollz/progressbar/v3"
)

var patternMapTitle = regexp.MustCompile(`(\w+)\-to\-(\w+) map:`)

type AlmanacMapEntry struct {
	DestinationStart, SourceStart, RangeLength int64
}
type AlmanacMap struct {
	Source      string
	Destination string
	Entries     []AlmanacMapEntry
}

func ParseSeeds(line string) []int64 {
	return fun.FieldsToInt64(strings.TrimPrefix(line, "seeds:"))
}

func ParseMap(lines []string) *AlmanacMap {
	matches := patternMapTitle.FindStringSubmatch(lines[0])
	if len(matches) == 0 {
		panic("wrong input")
	}
	return &AlmanacMap{
		Source:      matches[1],
		Destination: matches[2],
		Entries: fun.Map(lines[1:], func(line string) AlmanacMapEntry {
			vals := fun.FieldsToInt64(line)
			return AlmanacMapEntry{vals[0], vals[1], vals[2]}
		}),
	}
}

func ParseMaps(lines []string) []*AlmanacMap {
	return fun.Map(fun.Split(lines, func(v string) bool {
		return v == ""
	}), ParseMap)
}

func (m *AlmanacMap) MappedNumber(num int64) int64 {
	mapped := slices.IndexFunc(m.Entries, func(ame AlmanacMapEntry) bool {
		return num >= ame.SourceStart && num < ame.SourceStart+ame.RangeLength
	})
	if mapped == -1 {
		return num
	}
	entry := m.Entries[mapped]
	return entry.DestinationStart + num - entry.SourceStart
}

func GetLocation(num int64, almanac []*AlmanacMap) int64 {
	return fun.Reduce(almanac, func(m *AlmanacMap, prev int64) int64 {
		return m.MappedNumber(prev)
	}, num)
}

type day5 struct{}

func (day5) Part1(lines []string) any {
	seedNums := ParseSeeds(lines[0])
	almanac := ParseMaps(lines[2:])
	locations := fun.Map(seedNums, func(num int64) int64 {
		return GetLocation(num, almanac)
	})
	return fun.Min(locations)
}

func (day5) Part2(lines []string) any {
	seedNums := ParseSeeds(lines[0])
	almanac := ParseMaps(lines[2:])

	var wg sync.WaitGroup
	results := make([]int64, len(seedNums)/2)
	var resultsMtx sync.Mutex

	bar := progressbar.NewOptions(len(seedNums)/2,
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetElapsedTime(false),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionShowCount())

	for i := 0; i < len(seedNums); i += 2 {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			begin, length := seedNums[i], seedNums[i+1]
			min := int64(math.MaxInt64)
			for j := begin; j < begin+length; j++ {
				loc := GetLocation(j, almanac)
				if loc < min {
					min = loc
				}
			}

			resultsMtx.Lock()
			defer resultsMtx.Unlock()
			results[i/2] = min
			bar.Add(1)
		}()
	}

	wg.Wait()
	bar.Finish()

	return fun.Min(results)
}

func main() {
	aoc.Program(&day5{})
}
