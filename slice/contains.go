package slice

import (
	"slices"
)

// ContainsAny returns true if one needle is present in the haystack.
func ContainsAny[T comparable](haystack []T, needles []T) bool {
	for _, needle := range needles {
		if slices.Contains(haystack, needle) {
			return true
		}
	}
	return false
}

// ContainsAll returns true if all needles are present in the haystack.
func ContainsAll[T comparable](haystack []T, needles []T) bool {
	for _, needle := range needles {
		if !slices.Contains(haystack, needle) {
			return false
		}
	}
	return true
}
