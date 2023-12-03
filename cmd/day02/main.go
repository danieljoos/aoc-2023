package main

import (
	"strconv"
	"strings"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
)

type CubeSet struct {
	Red, Green, Blue int64
}

type CubeSets []CubeSet

type Game struct {
	ID   int64
	Sets CubeSets
}

type Games = []Game

func ParseCubeSet(str string) (result CubeSet) {
	splitColors := strings.Split(str, ",")
	for _, colorStr := range splitColors {
		splitColor := strings.Split(strings.TrimSpace(colorStr), " ")
		numStr, nameStr := splitColor[0], splitColor[1]
		val, err := strconv.ParseInt(numStr, 10, 8)
		if err != nil {
			panic(err)
		}
		switch nameStr {
		case "red":
			result.Red += val
		case "green":
			result.Green += val
		case "blue":
			result.Blue += val
		}
	}
	return result
}

func ParseLine(line string) (result Game) {
	splitGame := strings.Split(line, ":")
	idStr, setsStr := splitGame[0], splitGame[1]
	var err error
	if result.ID, err = strconv.ParseInt(strings.TrimPrefix(idStr, "Game "), 10, 8); err != nil {
		panic(err)
	}
	result.Sets = fun.Map(strings.Split(setsStr, ";"), ParseCubeSet)
	return result
}

func ParseInput(lines []string) (result Games) {
	return fun.Map(lines, ParseLine)
}

func (g Game) IsPossible(red, green, blue int64) bool {
	return fun.Every(g.Sets, func(v CubeSet) bool {
		return v.Red <= red && v.Green <= green && v.Blue <= blue
	})
}

func (g Game) MinSet() CubeSet {
	maxR := fun.Max(fun.Map(g.Sets, func(s CubeSet) int64 { return s.Red }))
	maxG := fun.Max(fun.Map(g.Sets, func(s CubeSet) int64 { return s.Green }))
	maxB := fun.Max(fun.Map(g.Sets, func(s CubeSet) int64 { return s.Blue }))
	return CubeSet{maxR, maxG, maxB}
}

func (cs CubeSet) Power() int64 {
	return cs.Red * cs.Green * cs.Blue
}

type day2 struct{}

func (*day2) Part1(lines []string) any {
	games := ParseInput(lines)
	possibleGames := fun.Filter(games, func(g Game) bool { return g.IsPossible(12, 13, 14) })
	return fun.Sum(fun.Map(possibleGames, func(g Game) int64 { return g.ID }))
}

func (*day2) Part2(lines []string) any {
	games := ParseInput(lines)
	minSets := fun.Map(games, func(g Game) CubeSet { return g.MinSet() })
	return fun.Sum(fun.Map(minSets, func(cs CubeSet) int64 { return cs.Power() }))
}

func main() {
	aoc.Program(&day2{})
}
