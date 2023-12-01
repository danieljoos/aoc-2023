package main

import (
	"regexp"
	"strings"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
)

var patternSpelled = regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func fromByte(b byte) int {
	return int(b - '0')
}

func fromString(str string) int {
	switch str {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return 0
}

func findAllOverlapIndex(pattern regexp.Regexp, str string) (results [][]int) {
	byteStr := []byte(str)
	moved := 0
	for {
		if len(byteStr) == 0 {
			break
		}
		p := pattern.FindIndex(byteStr)
		if p == nil {
			break
		}
		results = append(results, []int{moved + p[0], moved + p[1]})
		byteStr = byteStr[p[0]+1:]
		moved += p[0] + 1
	}
	return results
}

type day1 struct{}

var _ aoc.Day = (*day1)(nil)

func (*day1) Part1(lines []string) any {
	return fun.Sum(fun.Map(lines, func(line string) int {
		i1 := line[strings.IndexFunc(line, isNum)]
		i2 := line[strings.LastIndexFunc(line, isNum)]
		return fromByte(i1)*10 + fromByte(i2)
	}))
}

func (*day1) Part2(lines []string) any {
	return fun.Sum(fun.Map(lines, func(line string) int {
		pos := findAllOverlapIndex(*patternSpelled, line)
		i1 := strings.IndexFunc(line, isNum)
		i2 := strings.LastIndexFunc(line, isNum)

		// First digit
		v1 := 0
		if len(pos) > 0 && (pos[0][0] < i1 || i1 == -1) {
			v1 = fromString(line[pos[0][0]:pos[0][1]])
		} else {
			v1 = fromByte(line[i1])
		}

		// Last digit
		v2 := 0
		if last := len(pos) - 1; len(pos) > 0 && pos[last][0] > i2 {
			v2 = fromString(line[pos[last][0]:pos[last][1]])
		} else {
			v2 = fromByte(line[i2])
		}

		return v1*10 + v2
	}))
}

func main() {
	aoc.Program(&day1{})
}
