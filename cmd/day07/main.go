package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
)

type HandType int

const (
	HandTypeHighCard HandType = iota
	HandTypeOnePair
	HandTypeTwoPair
	HandTypeThree
	HandTypeFullHouse
	HandTypeFour
	HandTypeFive
)

type Strength struct {
	value string
	count int
}

type Hand struct {
	orig      string
	bid       int
	typ       HandType
	strengths []Strength
}

type Comparator string

func (c Comparator) CardValue(card string) int {
	return strings.Index(string(c), card)
}

func (c Comparator) CompareStrengths(lhs, rhs Strength) int {
	if lhs.count > rhs.count {
		return -1
	}
	if rhs.count > lhs.count {
		return 1
	}
	vl := c.CardValue(lhs.value)
	vr := c.CardValue(rhs.value)
	if vl > vr {
		return -1
	}
	if vr > vl {
		return 1
	}
	return 0
}

func (c Comparator) CompareHands(lhs, rhs *Hand) int {
	if lhs.typ > rhs.typ {
		return -1
	}
	if rhs.typ > lhs.typ {
		return 1
	}
	for i := 0; i < 5; i++ {
		lv := c.CardValue(string(lhs.orig[i]))
		rv := c.CardValue(string(rhs.orig[i]))
		if lv > rv {
			return -1
		}
		if rv > lv {
			return 1
		}
	}
	return 0
}

type Logic interface {
	Comparator() Comparator
	ParseStrengths(str string) []Strength
}

type LogicPart1 struct{}

func (LogicPart1) Comparator() Comparator {
	return "23456789TJQKA"
}

func (p LogicPart1) ParseStrengths(str string) []Strength {
	strengths := []Strength{}
	for i := 0; i < len(str); i++ {
		c := string(str[i])
		count := strings.Count(str, c)
		str := Strength{value: c, count: count}
		if idx := slices.IndexFunc(strengths, func(s Strength) bool {
			return s.count == count && s.value == c
		}); idx >= 0 {
			continue
		}
		strengths = append(strengths, str)
	}
	slices.SortFunc(strengths, p.Comparator().CompareStrengths)
	return strengths
}

type LogicPart2 struct{}

func (LogicPart2) Comparator() Comparator {
	return "J23456789TQKA"
}

func (p LogicPart2) ParseStrengths(str string) []Strength {
	strengths := []Strength{}
	for i := 0; i < len(str); i++ {
		c := string(str[i])
		count := strings.Count(str, c)
		str := Strength{value: c, count: count}
		if idx := slices.IndexFunc(strengths, func(s Strength) bool {
			return s.count == count && s.value == c
		}); idx >= 0 {
			continue
		}
		strengths = append(strengths, str)
	}
	slices.SortFunc(strengths, p.Comparator().CompareStrengths)
	// Handle jokers:
	if strengths[0].value != "J" {
		strengths[0].count += strings.Count(str, "J")
		strengths = fun.Filter(strengths, func(v Strength) bool { return v.value != "J" })
	} else if len(strengths) > 1 {
		strengths[1].count += strings.Count(str, "J")
		strengths = fun.Filter(strengths, func(v Strength) bool { return v.value != "J" })
	}
	return strengths
}

func ParseHand(line string, l Logic) *Hand {
	fields := strings.Fields(line)
	bid, _ := strconv.Atoi(fields[1])
	handStr := fields[0]
	strengths := l.ParseStrengths(handStr)

	// Deduce "type" of the hand
	typ := HandTypeHighCard
strloop:
	for _, s := range strengths {
		switch s.count {
		case 2:
			if typ == HandTypeThree {
				typ = HandTypeFullHouse
				break strloop
			}
			typ++
		case 3:
			typ = HandTypeThree
		case 4:
			typ = HandTypeFour
			break strloop
		case 5:
			typ = HandTypeFive
			break strloop
		}
	}

	return &Hand{
		orig:      handStr,
		bid:       bid,
		typ:       typ,
		strengths: strengths,
	}
}

func TotalWinnings(lines []string, l Logic) int {
	hands := fun.Map(lines, func(line string) *Hand { return ParseHand(line, l) })
	slices.SortFunc(hands, l.Comparator().CompareHands)
	slices.Reverse(hands)
	result := 0
	for i, h := range hands {
		rank := i + 1
		result += h.bid * rank
	}
	return result
}

type day7 struct{}

func (day7) Part1(lines []string) any {
	return TotalWinnings(lines, &LogicPart1{})
}

func (day7) Part2(lines []string) any {
	return TotalWinnings(lines, &LogicPart2{})
}

func main() {
	aoc.Program(day7{})
}
