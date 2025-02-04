package gotils

import (
	"slices"
)

// SliceContainsAny returns true if one needle is present in the haystack.
func SliceContainsAny[T comparable](haystack []T, needles []T) bool {
	for _, needle := range needles {
		if slices.Contains(haystack, needle) {
			return true
		}
	}
	return false
}

// SliceContainsAll returns true if all needles are present in the haystack.
func SliceContainsAll[T comparable](haystack []T, needles []T) bool {
	for _, needle := range needles {
		if !slices.Contains(haystack, needle) {
			return false
		}
	}
	return true
}

// SliceContainsNone returns true if no needles are present in the haystack.
func SliceContainsNone[T comparable](haystack []T, needles []T) bool {
	for _, needle := range needles {
		if slices.Contains(haystack, needle) {
			return false
		}
	}
	return true
}
