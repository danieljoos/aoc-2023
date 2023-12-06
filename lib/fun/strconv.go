package fun

import (
	"strconv"
	"strings"
)

func FieldsToInt64(input string) []int64 {
	return Map(strings.Fields(input), func(f string) int64 {
		v, err := strconv.ParseInt(f, 10, 64)
		if err != nil {
			panic(err)
		}
		return v
	})
}
