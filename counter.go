package gotils

import (
	"sync"
)

// Counter is a thread safe way to count comparable things.
type Counter[T comparable] struct {
	values map[T]uint
	m      sync.Mutex
}

// NewCounter creates a new Counter.
func NewCounter[T comparable](keys []T) *Counter[T] {
	c := Counter[T]{
		values: map[T]uint{},
	}

	for _, key := range keys {
		c.values[key] = 0
	}

	return &c
}

// AddKey adds a key if it doesn't exist already.
func (c *Counter[T]) AddKey(key T) {
	c.m.Lock()
	defer c.m.Unlock()

	if _, ok := c.values[key]; !ok {
		c.values[key] = 0
	}
}

// KeyExists checks if a key exists.
func (c *Counter[T]) KeyExists(key T) bool {
	c.m.Lock()
	defer c.m.Unlock()

	_, ok := c.values[key]
	return ok
}

// Increment adds 1 to the count. If the key doesn't exist, it's created.
func (c *Counter[T]) Increment(key T) {
	c.m.Lock()
	defer c.m.Unlock()

	c.values[key] = c.values[key] + 1
}

// IncrementBy adds the value to the count. If the key doesn't exist, it's created.
func (c *Counter[T]) IncrementBy(key T, value uint) {
	c.m.Lock()
	defer c.m.Unlock()

	c.values[key] = c.values[key] + value
}

// Increment adds 1 to the count if the key exists and return true.
func (c *Counter[T]) IncrementIfKeyExists(key T) bool {
	c.m.Lock()
	defer c.m.Unlock()

	if n, ok := c.values[key]; ok {
		c.values[key] = n + 1
		return true
	}

	return false
}

// Increment adds the value to the count if the key exists and return true.
func (c *Counter[T]) IncrementByIfKeyExists(key T, value uint) bool {
	c.m.Lock()
	defer c.m.Unlock()

	if n, ok := c.values[key]; ok {
		c.values[key] = n + value
		return true
	}

	return false
}

// Values returns the map of counts
func (c *Counter[T]) Values() map[T]uint {
	c.m.Lock()
	defer c.m.Unlock()

	return c.values
}

// SetToZero sets all key counts to zero
func (c *Counter[T]) SetToZero() {
	c.m.Lock()
	defer c.m.Unlock()

	for k := range c.values {
		c.values[k] = 0
	}
}

// Clear removes all keys
func (c *Counter[T]) Clear() {
	c.m.Lock()
	defer c.m.Unlock()

	c.values = map[T]uint{}
}
