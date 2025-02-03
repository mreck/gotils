package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	assert.Equal(t, Unique([]int{1, 2, 3, 2}), []int{1, 2, 3})
	assert.Equal(t, Unique([]int{1, 2, 3, 1, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, Unique([]int{3, 2, 1, 3, 2, 1}), []int{3, 2, 1})
}

func TestSortedUnique(t *testing.T) {
	assert.Equal(t, SortedUnique([]int{1, 2, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, SortedUnique([]int{1, 2, 2, 3, 2, 2}), []int{1, 2, 3, 2})
}

func TestSortAndMakeUnique(t *testing.T) {
	assert.Equal(t, SortAndMakeUnique([]int{1, 2, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, SortAndMakeUnique([]int{1, 2, 2, 3, 2, 2}), []int{1, 2, 3})
}
