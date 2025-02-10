package gotils_test

import (
	"testing"

	"github.com/mreck/gotils"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, gotils.Min(1, 2))
	assert.Equal(t, 1, gotils.Min(2, 1))

	assert.Equal(t, "a", gotils.Min("a", "b"))
	assert.Equal(t, "a", gotils.Min("b", "a"))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, gotils.Max(1, 2))
	assert.Equal(t, 2, gotils.Max(2, 1))

	assert.Equal(t, "b", gotils.Max("a", "b"))
	assert.Equal(t, "b", gotils.Max("b", "a"))
}
