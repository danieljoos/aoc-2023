package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
)

type day4 struct{}

type Card struct {
	ID             int64
	WinningNumbers []int64
	MyNumbers      []int64
}

func ParseNumbers(data string) []int64 {
	result := []int64{}
	for _, s := range strings.Fields(data) {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, v)
	}
	return result
}

func ParseCard(line string) *Card {
	result := &Card{}
	var err error

	splitCard := strings.Split(line, ":")
	cardIDStr, dataStr := splitCard[0], splitCard[1]
	if result.ID, err = strconv.ParseInt(strings.Fields(cardIDStr)[1], 10, 64); err != nil {
		panic(err)
	}

	splitData := strings.Split(dataStr, "|")
	result.WinningNumbers = ParseNumbers(splitData[0])
	result.MyNumbers = ParseNumbers(splitData[1])

	return result
}

func (day4) Part1(lines []string) any {
	cards := fun.Map(lines, ParseCard)
	return fun.Sum(fun.Map(cards, func(c *Card) int64 {
		// Linear, but good enough for this task
		myWinningNums := fun.Filter(c.WinningNumbers, func(v int64) bool {
			return slices.Contains(c.MyNumbers, v)
		})
		return fun.Reduce(myWinningNums, func(v, prev int64) int64 {
			if prev == 0 {
				return 1
			}
			return prev * 2
		}, 0)
	}))
}

func (day4) Part2(lines []string) any {
	cards := fun.Map(lines, ParseCard)
	copyCount := map[int64]int64{}
	return fun.Sum(fun.Map(cards, func(c *Card) int64 {
		myWinningNums := fun.Filter(c.WinningNumbers, func(v int64) bool {
			return slices.Contains(c.MyNumbers, v)
		})
		copies := copyCount[c.ID] + 1
		for j := 1; j <= len(myWinningNums); j++ {
			copyCount[c.ID+int64(j)] += copies
		}
		return copies
	}))
}

func main() {
	aoc.Program(&day4{})
}
