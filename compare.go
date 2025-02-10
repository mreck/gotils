package gotils

import (
	"golang.org/x/exp/constraints"
)

// Min returns the smaller value
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger value
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
