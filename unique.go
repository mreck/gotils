package gotils

// MakeSliceUnique returns a new slice containing the first appearance of each value.
func MakeSliceUnique[T comparable](array *[]T) {
	if len(*array) == 0 {
		return
	}

	seen := map[T]struct{}{}
	seen[(*array)[0]] = struct{}{}

	next := 1
	for i := 1; i < len(*array); i++ {
		v := (*array)[i]
		if _, ok := seen[v]; !ok {
			(*array)[next] = v
			next = next + 1
			seen[v] = struct{}{}
		}
	}

	*array = (*array)[0:next]
}

// MakeSortedSliceUnique returns a new slice containing the first appearance of each value.
func MakeSortedSliceUnique[T comparable](array *[]T) {
	if len(*array) == 0 {
		return
	}

	next := 1
	for i := 1; i < len(*array); i++ {
		v := (*array)[i]
		p := (*array)[next-1]
		if v != p {
			(*array)[next] = v
			next = next + 1
		}
	}

	*array = (*array)[0:next]
}
