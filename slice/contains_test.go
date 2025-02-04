package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsAny(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	assert.True(t, ContainsAny(data, []int{4, 2, 0}))
	assert.True(t, ContainsAny(data, []int{0, 7}))

	assert.False(t, ContainsAny(data, []int{0, 9}))
}

func TestContainsAll(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	assert.True(t, ContainsAll(data, []int{4, 2}))
	assert.True(t, ContainsAll(data, []int{7}))

	assert.False(t, ContainsAll(data, []int{4, 2, 0}))
}

func TestContainsNone(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	assert.True(t, ContainsNone(data, []int{8, 0}))
	assert.True(t, ContainsNone(data, []int{9}))

	assert.False(t, ContainsNone(data, []int{4}))
	assert.False(t, ContainsNone(data, []int{0, 4}))
}
