package gotils

import (
	"maps"
)

func GetMapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func GetMapValues[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	m2 := map[K]V{}
	maps.Copy(m2, m)
	return m2
}
