package gotils_test

import (
	"testing"

	"github.com/mreck/gotils"
	"github.com/stretchr/testify/assert"
)

func TestIsSet(t *testing.T) {
	assert.Equal(t, true, gotils.NewMaybe(int(1)).IsSet())
	assert.Equal(t, false, gotils.EmptyMaybe[int]().IsSet())
}

func TestGet(t *testing.T) {
	m := gotils.EmptyMaybe[int]()
	v, ok := m.Get()
	assert.Equal(t, false, ok)
	assert.Equal(t, int(0), v)

	m = gotils.NewMaybe(int(1))
	v, ok = m.Get()
	assert.Equal(t, true, ok)
	assert.Equal(t, int(1), v)
}

func TestSet(t *testing.T) {
	var m gotils.Maybe[int]
	v, ok := m.Get()
	assert.Equal(t, false, ok)
	assert.Equal(t, int(0), v)

	m.Set(1)
	v, ok = m.Get()
	assert.Equal(t, true, ok)
	assert.Equal(t, int(1), v)
}

func TestUnset(t *testing.T) {
	m := gotils.NewMaybe(int(1))
	v, ok := m.Get()
	assert.Equal(t, true, ok)
	assert.Equal(t, int(1), v)

	m.Unset()
	v, ok = m.Get()
	assert.Equal(t, false, ok)
	assert.Equal(t, int(0), v)
}
