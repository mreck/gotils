package gotils_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mreck/gotils"
)

func TestMakeSliceUnique(t *testing.T) {
	var d []int

	d = []int{1, 2, 3}
	gotils.MakeSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 3, 2}
	gotils.MakeSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 3, 1, 2, 3}
	gotils.MakeSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{3, 2, 1, 3, 2, 1}
	gotils.MakeSliceUnique(&d)
	assert.Equal(t, []int{3, 2, 1}, d)
}

func TestMakeSliceUniqueFunc(t *testing.T) {
	var d []int

	hasher := func(v int) string { return fmt.Sprintf("%d", v) }

	d = []int{1, 2, 3}
	gotils.MakeSliceUniqueFunc(&d, hasher)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 3, 2}
	gotils.MakeSliceUniqueFunc(&d, hasher)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 3, 1, 2, 3}
	gotils.MakeSliceUniqueFunc(&d, hasher)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{3, 2, 1, 3, 2, 1}
	gotils.MakeSliceUniqueFunc(&d, hasher)
	assert.Equal(t, []int{3, 2, 1}, d)
}

func TestMakeSortedSliceUnique(t *testing.T) {
	var d []int

	d = []int{1, 2, 3}
	gotils.MakeSortedSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 2, 3}
	gotils.MakeSortedSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3}, d)

	d = []int{1, 2, 2, 3, 2, 2}
	gotils.MakeSortedSliceUnique(&d)
	assert.Equal(t, []int{1, 2, 3, 2}, d)
}

func TestMakeSortedSliceUniqueFunc(t *testing.T) {
	var d []string

	d = []string{"a", "A", "b"}
	gotils.MakeSortedSliceUniqueFunc(&d, strings.EqualFold)
	assert.Equal(t, []string{"a", "b"}, d)

	d = []string{"A", "a", "b"}
	gotils.MakeSortedSliceUniqueFunc(&d, strings.EqualFold)
	assert.Equal(t, []string{"A", "b"}, d)
}
