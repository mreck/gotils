package gotils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mreck/gotils"
)

func TestSliceContainsAny(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	assert.True(t, gotils.SliceContainsAny(data, []int{4, 2, 0}))
	assert.True(t, gotils.SliceContainsAny(data, []int{0, 7}))

	assert.False(t, gotils.SliceContainsAny(data, []int{0, 9}))
}

func TestSliceContainsAll(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	assert.True(t, gotils.SliceContainsAll(data, []int{4, 2}))
	assert.True(t, gotils.SliceContainsAll(data, []int{7}))

	assert.False(t, gotils.SliceContainsAll(data, []int{4, 2, 0}))
}

func TestSliceContainsNone(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	assert.True(t, gotils.SliceContainsNone(data, []int{8, 0}))
	assert.True(t, gotils.SliceContainsNone(data, []int{9}))

	assert.False(t, gotils.SliceContainsNone(data, []int{4}))
	assert.False(t, gotils.SliceContainsNone(data, []int{0, 4}))
}
