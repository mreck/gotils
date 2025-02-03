package slice

import (
	"cmp"
	"slices"
)

// Unique returns a new slice containing the first appearance of each value.
func Unique[T comparable](array []T) []T {
	found := map[T]struct{}{}
	result := make([]T, 0, len(array))

	for _, v := range array {
		if _, ok := found[v]; !ok {
			result = append(result, v)
			found[v] = struct{}{}
		}
	}

	return result
}

// SortedUnique returns a new slice containing the first appearance of each value.
// Only works for sorted slices.
func SortedUnique[T comparable](array []T) []T {
	result := make([]T, 0, len(array))

	if len(array) > 0 {
		result = append(result, array[0])
	}

	j := 0
	for i := 1; i < len(array); i++ {
		v := array[i]
		if result[j] != v {
			result = append(result, v)
			j += 1
		}
	}

	return result
}

// SortAndMakeUnique returns a new slice with sorted unique values.
func SortAndMakeUnique[T cmp.Ordered](array []T) []T {
	result := Clone(array)
	slices.Sort(result)

	j := 0
	for i := 1; i < len(result); i++ {
		if result[j] != result[i] {
			j += 1
			result[j] = result[i]
		}
	}
	result = result[0 : j+1]

	return result
}
