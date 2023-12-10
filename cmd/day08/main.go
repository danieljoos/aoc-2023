package main

import (
	"regexp"
	"strings"

	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
	"golang.org/x/exp/maps"

	"github.com/TheAlgorithms/Go/math/lcm"
)

var patternNode = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

type Node struct {
	value, left, right string
}

type NodeMap = map[string]*Node

func ParseNodes(lines []string) NodeMap {
	return fun.Reduce(fun.Map(lines, func(line string) *Node {
		m := patternNode.FindStringSubmatch(line)
		return &Node{m[1], m[2], m[3]}
	}), func(n *Node, prev NodeMap) NodeMap {
		prev[n.value] = n
		return prev
	}, NodeMap{})
}

type day8 struct{}

func (day8) Part1(lines []string) any {
	instructions := strings.Split(lines[0], "")
	nodes := ParseNodes(lines[2:])
	curr := "AAA"

	i := 0
	steps := 0
	for {
		n := nodes[curr]
		if n.value == "ZZZ" {
			break
		}

		instr := instructions[i]
		i++
		if i >= len(instructions) {
			i = 0
		}

		switch instr {
		case "L":
			curr = n.left
		case "R":
			curr = n.right
		default:
			panic("invalid instruction")
		}
		steps++
	}
	return steps
}

func (day8) Part2(lines []string) any {
	instructions := strings.Split(lines[0], "")
	nodes := ParseNodes(lines[2:])
	startNodes := fun.Filter(maps.Keys(nodes), func(v string) bool { return strings.HasSuffix(v, "A") })

	result := int64(1)
	for _, curr := range startNodes {
		steps := int64(0)
		i := 0
		for {
			n := nodes[curr]
			if strings.HasSuffix(n.value, "Z") {
				break
			}

			instr := instructions[i]
			i++
			if i >= len(instructions) {
				i = 0
			}

			switch instr {
			case "L":
				curr = n.left
			case "R":
				curr = n.right
			default:
				panic("invalid instruction")
			}
			steps++
		}
		result = lcm.Lcm(result, steps)
	}
	return result
}

func main() {
	aoc.Program(&day8{})
}
