package gotils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCounter(t *testing.T) {
	assert.Equal(t, NewCounter[string](nil), &Counter[string]{
		values: map[string]uint{},
	})

	assert.Equal(t, NewCounter([]string{"foo", "bar"}), &Counter[string]{
		values: map[string]uint{"foo": 0, "bar": 0},
	})
}

func TestAddKey(t *testing.T) {
	c := NewCounter([]string{"foo"})
	c.AddKey("bar")
	c.AddKey("baz")

	assert.Equal(t, c, &Counter[string]{
		values: map[string]uint{"foo": 0, "bar": 0, "baz": 0},
	})
}

func TestKeyExists(t *testing.T) {
	c := NewCounter([]string{"foo", "bar"})

	assert.True(t, c.KeyExists("foo"))
	assert.True(t, c.KeyExists("bar"))
	assert.False(t, c.KeyExists("baz"))
}

func TestKeyExistsFunc(t *testing.T) {
	c := NewCounter([]string{"foo", "bar"})

	assert.True(t, c.KeyExistsFunc(func(key string) bool { return key == "foo" }))
	assert.True(t, c.KeyExistsFunc(func(key string) bool { return key == "bar" }))
	assert.False(t, c.KeyExistsFunc(func(key string) bool { return key == "baz" }))
}

func TestIncrement(t *testing.T) {
	c := NewCounter([]string{"foo", "bar"})

	assert.Equal(t, c, &Counter[string]{
		values: map[string]uint{"foo": 0, "bar": 0},
	})

	c.Increment("foo")

	assert.Equal(t, c, &Counter[string]{
		values: map[string]uint{"foo": 1, "bar": 0},
	})

	c.Increment("foo")
	c.Increment("bar")

	assert.Equal(t, c, &Counter[string]{
		values: map[string]uint{"foo": 2, "bar": 1},
	})
}

func TestIncrementBy(t *testing.T) {
	c := NewCounter([]string{"foo", "bar"})

	assert.Equal(t, c, &Counter[string]{
		values: map[string]uint{"foo": 0, "bar": 0},
	})

	c.IncrementBy("foo", 2)

	assert.Equal(t, c, &Counter[string]{
		values: map[string]uint{"foo": 2, "bar": 0},
	})

	c.IncrementBy("foo", 2)
	c.IncrementBy("bar", 2)

	assert.Equal(t, c, &Counter[string]{
		values: map[string]uint{"foo": 4, "bar": 2},
	})
}

func TestIncrementIfKeyExists(t *testing.T) {
	c := NewCounter([]string{"foo"})

	assert.True(t, c.IncrementIfKeyExists("foo"))
	assert.True(t, c.IncrementIfKeyExists("foo"))
	assert.False(t, c.IncrementIfKeyExists("bar"))
	assert.False(t, c.IncrementIfKeyExists("bar"))

	assert.Equal(t, c, &Counter[string]{
		values: map[string]uint{"foo": 2},
	})
}

func TestIncremenBytIfKeyExists(t *testing.T) {
	c := NewCounter([]string{"foo"})

	assert.True(t, c.IncrementByIfKeyExists("foo", 2))
	assert.True(t, c.IncrementByIfKeyExists("foo", 3))
	assert.False(t, c.IncrementByIfKeyExists("bar", 7))
	assert.False(t, c.IncrementByIfKeyExists("bar", 8))

	assert.Equal(t, c, &Counter[string]{
		values: map[string]uint{"foo": 5},
	})
}

func TestIncrementFunc(t *testing.T) {
	c := NewCounter([]string{"foo", "bar"})

	assert.True(t, c.IncrementFunc(func(key string, cnt uint) (uint, bool) {
		if key == "foo" {
			return 1, true
		}
		return 0, true
	}))

	assert.True(t, c.IncrementFunc(func(key string, cnt uint) (uint, bool) {
		if key == "bar" {
			return 2, false
		}
		return 0, true
	}))

	assert.Equal(t, c.values, map[string]uint{"foo": 1, "bar": 2})
}

func TestKeys(t *testing.T) {
	c := NewCounter([]string{"foo", "bar"})

	keys := c.Keys()
	sort.Strings(keys)
	assert.Equal(t, keys, []string{"bar", "foo"})
}

func TestValues(t *testing.T) {
	c := NewCounter([]string{"foo", "bar"})
	c.Increment("foo")
	c.Increment("foo")
	c.Increment("bar")

	assert.Equal(t, c.Values(), map[string]uint{"foo": 2, "bar": 1})
}

func TestSetToZero(t *testing.T) {
	c := NewCounter([]string{"foo", "bar"})
	c.Increment("foo")
	c.Increment("foo")
	c.Increment("bar")
	c.SetToZero()

	assert.Equal(t, c.Values(), map[string]uint{"foo": 0, "bar": 0})
}

func TestClear(t *testing.T) {
	c := NewCounter([]string{"foo", "bar"})
	c.Increment("foo")
	c.Increment("foo")
	c.Increment("bar")
	c.Clear()

	assert.Equal(t, c.Values(), map[string]uint{})
}
