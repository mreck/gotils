package gotils

// Maybe wraps a value and adds validity
type Maybe[T any] struct {
	Value T
	Valid bool
}

// Result joins a value and an error together
type Result[T any] struct {
	Value T
	Err   error
}
