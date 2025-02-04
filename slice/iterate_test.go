package slice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t, Map([]int{1, 2, 3, 4}, func(n int) int {
		return n * n
	}), []int{1, 4, 9, 16})

	assert.Equal(t, Map([]int{1, 2, 3, 4}, func(n int) string {
		return strconv.FormatInt(int64(n), 10)
	}), []string{"1", "2", "3", "4"})
}

func TestReduce(t *testing.T) {
	assert.Equal(t, Reduce([]int{1, 2, 3, 4}, 0, func(acc int, n int) int {
		return acc + n
	}), 10)

	assert.Equal(t, Reduce([]int{1, 2, 3, 4}, "0", func(acc string, n int) string {
		return acc + strconv.FormatInt(int64(n), 10)
	}), "01234")
}

func TestClone(t *testing.T) {
	initial := []int{1, 2, 3, 4, 5}
	cloned := Clone(initial)

	assert.Equal(t, initial, cloned)
	assert.False(t, &initial == &cloned)
}
