package gotils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUniqueSlice(t *testing.T) {
	assert.Equal(t, CreateUniqueSlice([]int{1, 2, 3, 2}), []int{1, 2, 3})
	assert.Equal(t, CreateUniqueSlice([]int{1, 2, 3, 1, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, CreateUniqueSlice([]int{3, 2, 1, 3, 2, 1}), []int{3, 2, 1})
}

func TestCreateUniqueSliceFromSorted(t *testing.T) {
	assert.Equal(t, CreateUniqueSliceFromSorted([]int{1, 2, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, CreateUniqueSliceFromSorted([]int{1, 2, 2, 3, 2, 2}), []int{1, 2, 3, 2})
}

func TestSortAndMakeSliceUnique(t *testing.T) {
	assert.Equal(t, SortAndMakeSliceUnique([]int{1, 2, 2, 3}), []int{1, 2, 3})
	assert.Equal(t, SortAndMakeSliceUnique([]int{1, 2, 2, 3, 2, 2}), []int{1, 2, 3})
}
