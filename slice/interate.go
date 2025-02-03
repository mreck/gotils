package slice

// Map returns a new slice of mapped values.
func Map[T any, R any](array []T, fn func(T) R) []R {
	result := make([]R, 0, len(array))

	for _, v := range array {
		result = append(result, fn(v))
	}

	return result
}

// Clone returns a new slice with the same values.
func Clone[T any](array []T) []T {
	result := make([]T, 0, len(array))
	result = append(result, array...)
	return result
}
