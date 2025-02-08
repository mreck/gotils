package gotils

import (
	"fmt"
	"strings"
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

func TestMakeSliceUniqueFunc(t *testing.T) {
	var d []int

	hasher := func(v int) string { return fmt.Sprintf("%d", v) }

	d = []int{1, 2, 3}
	MakeSliceUniqueFunc(&d, hasher)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 3, 2}
	MakeSliceUniqueFunc(&d, hasher)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 3, 1, 2, 3}
	MakeSliceUniqueFunc(&d, hasher)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{3, 2, 1, 3, 2, 1}
	MakeSliceUniqueFunc(&d, hasher)
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

func TestMakeSortedSliceUniqueFunc(t *testing.T) {
	var d []string

	d = []string{"a", "A", "b"}
	MakeSortedSliceUniqueFunc(&d, strings.EqualFold)
	assert.Equal(t, []string{"a", "b"}, d)

	d = []string{"A", "a", "b"}
	MakeSortedSliceUniqueFunc(&d, strings.EqualFold)
	assert.Equal(t, []string{"A", "b"}, d)
}
