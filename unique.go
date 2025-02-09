package gotils

// MakeSliceUnique removes all but the first appearance of each value.
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
			next++
			seen[v] = struct{}{}
		}
	}

	*array = (*array)[0:next]
}

// MakeSliceUniqueFunc removes all but the first appearance of each value.
// Values are hashed using the hasher function.
func MakeSliceUniqueFunc[T any](array *[]T, hasher func(v T) string) {
	if len(*array) == 0 {
		return
	}

	seen := map[string]struct{}{}
	seen[hasher((*array)[0])] = struct{}{}

	next := 1
	for i := 1; i < len(*array); i++ {
		v := (*array)[i]
		h := hasher(v)
		if _, ok := seen[h]; !ok {
			(*array)[next] = v
			next++
			seen[h] = struct{}{}
		}
	}

	*array = (*array)[0:next]
}

// MakeSortedSliceUnique removes all but the first appearance of each value for pre-sorted slices.
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
			next++
		}
	}

	*array = (*array)[0:next]
}

// MakeSortedSliceUnique removes all but the first appearance of each value for pre-sorted slices.
// Values are compared using the equal function.
func MakeSortedSliceUniqueFunc[T any](array *[]T, equal func(a, b T) bool) {
	if len(*array) == 0 {
		return
	}

	next := 1
	for i := 1; i < len(*array); i++ {
		v := (*array)[i]
		p := (*array)[next-1]
		if !equal(v, p) {
			(*array)[next] = v
			next++
		}
	}

	*array = (*array)[0:next]
}
