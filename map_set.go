package gotils

import (
	"maps"
)

// MapDifference calculates (A \ B), based on the maps' keys
func MapDifference[K comparable, V any](a, b map[K]V) map[K]V {
	res := map[K]V{}
	for k, v := range a {
		if _, ok := b[k]; !ok {
			res[k] = v
		}
	}
	return res
}

// MapIntersection calculates (A ∩ B), based on the maps' keys, using A's values
func MapIntersection[K comparable, V any](a, b map[K]V) map[K]V {
	res := map[K]V{}
	for k, v := range a {
		if _, ok := b[k]; ok {
			res[k] = v
		}
	}
	return res
}

// MapSymetricDifference calculates ((A \ B) ∪ (B \ A)), based on the maps' keys
func MapSymetricDifference[K comparable, V any](a, b map[K]V) map[K]V {
	res := map[K]V{}
	for k, v := range a {
		if _, ok := b[k]; !ok {
			res[k] = v
		}
	}
	for k, v := range b {
		if _, ok := a[k]; !ok {
			res[k] = v
		}
	}
	return res
}

// MapUnion calculates (A ∪ B), based on the maps' keys, preferring A's values
func MapUnion[K comparable, V any](a, b map[K]V) map[K]V {
	res := map[K]V{}
	maps.Copy(res, a)
	for k, v := range b {
		if _, ok := a[k]; !ok {
			res[k] = v
		}
	}
	return res
}

// MapIsSubset calculates (A ⊆ B), based on the maps' keys
func MapIsSubset[K comparable, V any](a, b map[K]V) bool {
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}

// MapIsSuperset calculates (A ⊇ B), based on the maps' keys
func MapIsSuperset[K comparable, V any](a, b map[K]V) bool {
	for k := range b {
		if _, ok := a[k]; !ok {
			return false
		}
	}
	return true
}
