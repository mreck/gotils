package gotils

// MapSlice returns a new slice of mapped values.
func MapSlice[T any, R any](array []T, fn func(T) R) []R {
	result := make([]R, 0, len(array))

	for _, v := range array {
		result = append(result, fn(v))
	}

	return result
}

// ApplyToSlice applies all functions to the slice
func ApplyToSlice[T any](array []T, funcs ...func(T) T) {
	for _, fn := range funcs {
		for i, val := range array {
			array[i] = fn(val)
		}
	}
}

// ReduceSlice returns the reduced value.
func ReduceSlice[T any, R any](array []T, initial R, fn func(R, T) R) R {
	result := initial

	for _, v := range array {
		result = fn(result, v)
	}

	return result
}

// CloneSlice returns a new slice with the same values.
func CloneSlice[T any](array []T) []T {
	result := make([]T, 0, len(array))
	result = append(result, array...)
	return result
}

// FilterSlice removes all items from the slice that don't match the filter.
func FilterSlice[T any](array *[]T, filter func(T) bool) {
	next := 0

	for i := 0; i < len(*array); i++ {
		v := (*array)[i]
		if filter(v) {
			(*array)[next] = v
			next++
		}
	}

	*array = (*array)[:next]
}
