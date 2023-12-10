package main

import (
	"github.com/danieljoos/aoc-2023/lib/aoc"
	"github.com/danieljoos/aoc-2023/lib/fun"
)

func Extrapolate(nums []int64) (int64, int64) {
	if fun.Every(nums, func(n int64) bool { return n == 0 }) {
		return 0, 0
	}
	next := make([]int64, 0, len(nums)-1)
	for i, n := range nums[1:] {
		next = append(next, n-nums[i])
	}
	past, future := Extrapolate(next)
	return nums[0] - past, nums[len(nums)-1] + future
}

type day9 struct{}

func (day9) Part1(lines []string) any {
	return fun.Sum(fun.Map(lines, func(line string) int64 {
		_, x := Extrapolate(fun.FieldsToInt64(line))
		return x
	}))
}

func (day9) Part2(lines []string) any {
	return fun.Sum(fun.Map(lines, func(line string) int64 {
		x, _ := Extrapolate(fun.FieldsToInt64(line))
		return x
	}))
}

func main() {
	aoc.Program(&day9{})
}
