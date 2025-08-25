package gotils

import (
	"cmp"
)

func Clamp[T cmp.Ordered](lower, upper, value T) T {
	if value < lower {
		return lower
	}
	if value > upper {
		return upper
	}
	return value
}
