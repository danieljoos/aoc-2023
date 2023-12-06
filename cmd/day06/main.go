package main

import (
	"strconv"
	"strings"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
)

type Race struct {
	Duration int64
	Distance int64
}

func ParseRaces(lines []string) []Race {
	times := fun.FieldsToInt64(strings.TrimPrefix(lines[0], "Time:"))
	distances := fun.FieldsToInt64(strings.TrimPrefix(lines[1], "Distance:"))
	results := make([]Race, 0, len(times))
	for i := range times {
		results = append(results, Race{times[i], distances[i]})
	}
	return results
}

func ParseRaceCorrectKerning(lines []string) Race {
	time, _ := strconv.ParseInt(strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time:"), " ", ""), 10, 64)
	distance, _ := strconv.ParseInt(strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance:"), " ", ""), 10, 64)
	return Race{time, distance}
}

func WaysToWin(race Race) int {
	count := 0
	for i := int64(1); i < race.Duration; i++ {
		d := (race.Duration - i) * i
		if d > race.Distance {
			count++
		}
	}
	return count
}

type day6 struct{}

func (day6) Part1(lines []string) any {
	races := ParseRaces(lines)
	counts := fun.Map(races, WaysToWin)
	return fun.Product(counts)
}

func (day6) Part2(lines []string) any {
	race := ParseRaceCorrectKerning(lines)
	return WaysToWin(race)
}

func main() {
	aoc.Program(&day6{})
}
