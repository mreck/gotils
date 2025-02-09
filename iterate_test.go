package gotils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSlice(t *testing.T) {
	assert.Equal(t, MapSlice([]int{1, 2, 3, 4}, func(n int) int {
		return n * n
	}), []int{1, 4, 9, 16})

	assert.Equal(t, MapSlice([]int{1, 2, 3, 4}, func(n int) string {
		return strconv.FormatInt(int64(n), 10)
	}), []string{"1", "2", "3", "4"})
}

func TestApplyToSlice(t *testing.T) {
	data := []int{1, 2, 3, 4}
	ApplyToSlice(
		data,
		func(n int) int { return n + 1 },
		func(n int) int { return n * n })
	assert.Equal(t, data, []int{4, 9, 16, 25})
}

func TestReduceSlice(t *testing.T) {
	assert.Equal(t, ReduceSlice([]int{1, 2, 3, 4}, 0, func(acc int, n int) int {
		return acc + n
	}), 10)

	assert.Equal(t, ReduceSlice([]int{1, 2, 3, 4}, "0", func(acc string, n int) string {
		return acc + strconv.FormatInt(int64(n), 10)
	}), "01234")
}

func TestCloneSlice(t *testing.T) {
	initial := []int{1, 2, 3, 4, 5}
	cloned := CloneSlice(initial)

	assert.Equal(t, initial, cloned)
	assert.False(t, &initial == &cloned)
}

func TestFilterSlice(t *testing.T) {
	var d []int

	d = []int{1, 2, 3, 4, 5, 6}
	FilterSlice(&d, func(n int) bool { return (n % 2) == 0 })
	assert.Equal(t, []int{2, 4, 6}, d)

	d = []int{1, 2, 3, 4, 5, 6}
	FilterSlice(&d, func(n int) bool { return n >= 4 })
	assert.Equal(t, []int{4, 5, 6}, d)
}
