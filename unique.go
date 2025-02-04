package gotils

import (
	"cmp"
	"slices"
)

// CreateUniqueSlice returns a new slice containing the first appearance of each value.
func CreateUniqueSlice[T comparable](array []T) []T {
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

// CreateUniqueSliceFromSorted returns a new slice containing the first appearance of each value.
// Only works for sorted slices.
func CreateUniqueSliceFromSorted[T comparable](array []T) []T {
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

// SortAndMakeSliceUnique returns a new slice with sorted unique values.
func SortAndMakeSliceUnique[T cmp.Ordered](array []T) []T {
	result := CloneSlice(array)
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
