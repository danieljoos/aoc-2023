package aoc

import (
	"fmt"
	"os"

	"github.com/danieljoos/aoc-2023/lib/input"
)

type Day interface {
	Part1(lines []string) any
	Part2(lines []string) any
}

func Program(d Day) {
	if len(os.Args) < 3 {
		fmt.Println("usage: CMD part1|part2 <FILE>")
		os.Exit(1)
	}

	lines, err := input.ReadLines(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "part1":
		fmt.Println(d.Part1(lines))
	case "part2":
		fmt.Println(d.Part2(lines))
	default:
		fmt.Println("unknown part")
		os.Exit(1)
	}
}
