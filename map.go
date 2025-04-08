package gotils

import (
	"maps"
)

func GetMapKeys[K comparable, V any](src map[K]V) []K {
	r := make([]K, 0, len(src))
	for k := range src {
		r = append(r, k)
	}
	return r
}

func GetMapValues[K comparable, V any](src map[K]V) []V {
	r := make([]V, 0, len(src))
	for _, v := range src {
		r = append(r, v)
	}
	return r
}

func CopyMap[K comparable, V any](src map[K]V) map[K]V {
	dst := map[K]V{}
	maps.Copy(dst, src)
	return dst
}

func ExtendMap[K comparable, V any](dst map[K]V, srcs ...map[K]V) {
	for _, src := range srcs {
		for k, v := range src {
			dst[k] = v
		}
	}
}
