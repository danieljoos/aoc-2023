package main

import (
	"regexp"
	"strconv"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
)

var patternNums = regexp.MustCompile(`\d+`)

type NumPos = []int
type Part struct {
	Value int64
	Pos   NumPos
}
type Parts = []Part

func ParseNums(line string) [][]int {
	return patternNums.FindAllStringIndex(line, -1)
}

func IsSymbol(r byte) bool {
	return !(r >= '0' && r <= '9') && r != '.'
}

func IsAdjacentToSym(row int, numPos NumPos, lines []string) bool {
	leftMost, rightMost := numPos[0], numPos[1]
	isMin := leftMost == 0
	isMax := rightMost == len(lines[row])
	isMaxRow := row == len(lines)-1
	isMinRow := row == 0
	// Left of the number
	if !isMin && ((IsSymbol(lines[row][leftMost-1])) ||
		(!isMinRow && IsSymbol(lines[row-1][leftMost-1])) ||
		(!isMaxRow && IsSymbol(lines[row+1][leftMost-1]))) {
		return true
	}
	// Right of the number
	if !isMax && (IsSymbol(lines[row][rightMost]) ||
		(!isMinRow && IsSymbol(lines[row-1][rightMost])) ||
		(!isMaxRow && IsSymbol(lines[row+1][rightMost]))) {
		return true
	}
	for i := leftMost; i <= rightMost-1; i++ {
		if (!isMinRow && IsSymbol(lines[row-1][i])) ||
			(!isMaxRow && IsSymbol(lines[row+1][i])) {
			return true
		}
	}
	return false
}

type day3 struct{}

func (day3) Part1(lines []string) any {
	rows := fun.Map(lines, ParseNums)
	var sum int64 = 0
	for i, rowNums := range rows {
		adj := fun.Filter(rowNums, func(numPos NumPos) bool {
			return IsAdjacentToSym(i, numPos, lines)
		})
		rowPartNums := fun.Map(adj, func(numPos NumPos) int64 {
			s := lines[i][numPos[0]:numPos[1]]
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				panic(err)
			}
			return v
		})
		sum += fun.Sum(rowPartNums)
	}
	return sum
}

func (day3) Part2(lines []string) any {
	rows := fun.Map(lines, ParseNums)
	rowsParts := make([]Parts, 0, len(rows))
	for i, rowNums := range rows {
		adj := fun.Filter(rowNums, func(numPos NumPos) bool {
			return IsAdjacentToSym(i, numPos, lines)
		})
		parts := fun.Map(adj, func(numPos NumPos) Part {
			s := lines[i][numPos[0]:numPos[1]]
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				panic(err)
			}
			return Part{v, numPos}
		})
		rowsParts = append(rowsParts, parts)
	}
	sum := int64(0)
	for i, line := range lines {
		for j, c := range line {
			if c != '*' {
				continue
			}
			isAdjacentPart := func(p Part) bool {
				return p.Pos[0] <= j+1 && p.Pos[1] >= j
			}
			var adjParts Parts
			// Above
			if i > 0 {
				adjParts = append(adjParts, fun.Filter(rowsParts[i-1], isAdjacentPart)...)
			}
			// Below
			if i < len(lines)-1 {
				adjParts = append(adjParts, fun.Filter(rowsParts[i+1], isAdjacentPart)...)
			}
			// Same row
			if j > 0 {
				adjParts = append(adjParts, fun.Filter(rowsParts[i], isAdjacentPart)...)
			}
			if len(adjParts) == 2 {
				// It's a gear
				ratio := adjParts[0].Value * adjParts[1].Value
				sum += ratio
			}
		}
	}
	return sum
}

func main() {
	aoc.Program(&day3{})
}
