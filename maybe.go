package gotils

type Maybe[T any] struct {
	value T
	isSet bool
}

func EmptyMaybe[T any]() Maybe[T] {
	var m Maybe[T]
	return m
}

func NewMaybe[T any](value T) Maybe[T] {
	return Maybe[T]{value, true}
}

func (m Maybe[T]) IsSet() bool {
	return m.isSet
}

func (m Maybe[T]) Get() (T, bool) {
	return m.value, m.isSet
}

func (m *Maybe[T]) Set(value T) {
	m.value = value
	m.isSet = true
}

func (m *Maybe[T]) Unset() {
	var v T
	m.value = v
	m.isSet = false
}
