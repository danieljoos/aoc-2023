package main

import (
	"math"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
)

type Pixel bool
type Image [][]Pixel
type Galaxy struct{ X, Y int }
type ExpansionAreas struct {
	EmptyRows []int
	EmptyCols []int
}

func IsGalaxy(p Pixel) bool {
	return bool(p)
}

func IsEmpty(p Pixel) bool {
	return !bool(p)
}

func ParseImage(lines []string) Image {
	return fun.Map(lines, func(line string) []Pixel {
		row := make([]Pixel, 0, len(line))
		for _, b := range line {
			row = append(row, b == '#')
		}
		return row
	})
}

func FindExpansionAreas(universe Image) *ExpansionAreas {
	result := &ExpansionAreas{}
	for i := 0; i < len(universe); i++ {
		row := universe[i]
		if fun.Every(row, IsEmpty) {
			result.EmptyRows = append(result.EmptyRows, i)
		}
	}
	for i := 0; i < len(universe[0]); i++ {
		col := fun.Map(universe, func(row []Pixel) Pixel {
			return row[i]
		})
		if fun.Every(col, IsEmpty) {
			result.EmptyCols = append(result.EmptyCols, i)
		}
	}
	return result
}

func Galaxies(universe Image) []Galaxy {
	result := []Galaxy{}
	for y, row := range universe {
		for x, p := range row {
			if IsGalaxy(p) {
				result = append(result, Galaxy{x, y})
			}
		}
	}
	return result
}

func ExpansionEffects(g1, g2 Galaxy, expansionAreas *ExpansionAreas, factor int) (int, int) {
	minX := int(math.Min(float64(g1.X), float64(g2.X)))
	maxX := int(math.Max(float64(g1.X), float64(g2.X)))
	minY := int(math.Min(float64(g1.Y), float64(g2.Y)))
	maxY := int(math.Max(float64(g1.Y), float64(g2.Y)))

	expX := fun.Reduce(expansionAreas.EmptyCols, func(r int, prev int) int {
		if r > minX && r < maxX {
			prev++
		}
		return prev
	}, 0)
	expY := fun.Reduce(expansionAreas.EmptyRows, func(r int, prev int) int {
		if r > minY && r < maxY {
			prev++
		}
		return prev
	}, 0)

	return expX * (factor - 1), expY * (factor - 1)
}

func SumShortestPaths(galaxies []Galaxy, expansionAreas *ExpansionAreas, expansionFactor int) int {
	sum := 0
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			shortestPath := int(math.Abs(float64(g1.X)-float64(g2.X)) + math.Abs(float64(g1.Y)-float64(g2.Y)))
			sum += shortestPath
			expX, expY := ExpansionEffects(g1, g2, expansionAreas, expansionFactor)
			sum += expX + expY
		}
	}
	return sum
}

type day11 struct{}

func (day11) Part1(lines []string) any {
	universe := ParseImage(lines)
	return SumShortestPaths(Galaxies(universe), FindExpansionAreas(universe), 2)
}

func (day11) Part2(lines []string) any {
	universe := ParseImage(lines)
	return SumShortestPaths(Galaxies(universe), FindExpansionAreas(universe), 1000000)
}

func main() {
	aoc.Program(&day11{})
}
