package gotils

// Maybe wraps a value and adds validity
type Maybe[T any] struct {
	Value T
	Valid bool
}
