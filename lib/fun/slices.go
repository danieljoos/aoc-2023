package fun

import (
	"slices"

	"golang.org/x/exp/constraints"
)

func Map[T any, R any, S ~[]T](s S, pred func(v T) R) []R {
	result := make([]R, 0, len(s))
	for _, v := range s {
		result = append(result, pred(v))
	}
	return result
}

func Reduce[T any, R any, S ~[]T](s S, pred func(v T, prev R) R, init R) R {
	result := init
	for _, v := range s {
		result = pred(v, result)
	}
	return result
}

func Filter[T any, S ~[]T](s S, pred func(v T) bool) S {
	result := S{}
	for _, v := range s {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func Every[T any, S ~[]T](s S, pred func(v T) bool) bool {
	for _, v := range s {
		if !pred(v) {
			return false
		}
	}
	return true
}

func Sum[T constraints.Integer](s []T) T {
	return Reduce(s, func(v, prev T) T { return prev + v }, 0)
}

func Product[T constraints.Integer](s []T) T {
	return Reduce(s, func(v, prev T) T { return prev * v }, 1)
}

func Min[T constraints.Integer](s []T) T {
	return Reduce(s[1:], func(v, prev T) T {
		if v < prev {
			return v
		}
		return prev
	}, s[0])
}

func Max[T constraints.Integer](s []T) T {
	return Reduce(s[1:], func(v, prev T) T {
		if v > prev {
			return v
		}
		return prev
	}, s[0])
}

func Split[T any, S ~[]T](s S, pred func(v T) bool) []S {
	n := Count(s, pred) + 1
	result := make([]S, n)
	n--
	i := 0
	for i < n {
		idx := slices.IndexFunc(s, pred)
		if idx < 0 {
			break
		}
		result[i] = s[:idx]
		s = s[idx+1:]
		i++
	}
	result[i] = s
	return result
}

func Count[T any, S ~[]T](s S, pred func(v T) bool) int {
	return Reduce(s, func(v T, prev int) int {
		if pred(v) {
			return prev + 1
		}
		return prev
	}, 0)
}
