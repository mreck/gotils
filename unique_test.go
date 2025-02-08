package gotils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeSliceUnique(t *testing.T) {
	var d []int

	d = []int{1, 2, 3}
	MakeSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 3, 2}
	MakeSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 3, 1, 2, 3}
	MakeSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{3, 2, 1, 3, 2, 1}
	MakeSliceUnique(&d)
	assert.Equal(t, []int{3, 2, 1}, d)

}

func TestMakeSortedSliceUnique(t *testing.T) {
	var d []int

	d = []int{1, 2, 3}
	MakeSortedSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 2, 3}
	MakeSortedSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 2, 3, 2, 2}
	MakeSortedSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3, 2}, d)
}
