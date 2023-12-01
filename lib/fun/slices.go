package fun

func Map[T any, R any, S ~[]T](s S, pred func(v T) R) []R {
	result := make([]R, len(s))
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

func Sum(s []int) int {
	return Reduce(s, func(v, prev int) int { return prev + v }, 0)
}
