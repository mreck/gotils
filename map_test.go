package gotils_test

import (
	"sort"
	"testing"

	"github.com/mreck/gotils"
	"github.com/stretchr/testify/assert"
)

func TestGetMapKeys(t *testing.T) {
	m := map[int]int{1: 2, 3: 4, 5: 6}
	k := gotils.GetMapKeys(m)
	sort.Ints(k)
	assert.Equal(t, []int{1, 3, 5}, k)
}

func TestGetMapValues(t *testing.T) {
	m := map[int]int{1: 2, 3: 4, 5: 6}
	k := gotils.GetMapValues(m)
	sort.Ints(k)
	assert.Equal(t, []int{2, 4, 6}, k)
}

func TestCopyMap(t *testing.T) {
	m1 := map[int]int{1: 2, 3: 4, 5: 6}
	m2 := gotils.CopyMap(m1)
	assert.Equal(t, m1, m2)
	m1[1] = 7
	assert.Equal(t, 2, m2[1])
}
