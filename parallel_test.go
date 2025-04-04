package gotils_test

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mreck/gotils"
)

var (
	ErrParallelTest = errors.New("test")
)

func TestParallelFor(t *testing.T) {
	var out []int
	var errs []error
	var m sync.Mutex

	in := []int{1, 2, 3, 4, 5, 6}
	ctx := context.Background()

	errs = gotils.ParellelFor(ctx, in, 4, func(ctx context.Context, i int, n int) error {
		m.Lock()
		defer m.Unlock()
		out = append(out, n+1)
		return nil
	})

	sort.Ints(out)
	assert.Nil(t, errs)
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7}, out)

	errs = gotils.ParellelFor(ctx, in, 4, func(ctx context.Context, i int, n int) error {
		if n%2 == 0 {
			return fmt.Errorf("%w: %d", ErrParallelTest, i)
		}
		return nil
	})

	assert.Len(t, errs, 3)
	for _, err := range errs {
		assert.ErrorIs(t, err, ErrParallelTest)
	}
}

func TestParellelMap(t *testing.T) {
	var out []int
	var errs []error

	in := []int{1, 2, 3, 4, 5, 6}
	ctx := context.Background()

	out, errs = gotils.ParellelMap(ctx, in, 4, func(ctx context.Context, i int, n int) (int, error) {
		return n + 1, nil
	})

	assert.Equal(t, []int{2, 3, 4, 5, 6, 7}, out)
	assert.Nil(t, errs)

	out, errs = gotils.ParellelMap(ctx, in, 4, func(ctx context.Context, i int, n int) (int, error) {
		if n%2 == 0 {
			return 0, fmt.Errorf("%w: %d", ErrParallelTest, i)
		}
		return n, nil
	})

	assert.Equal(t, []int{1, 0, 3, 0, 5, 0}, out)
	assert.Len(t, errs, 3)
	for _, err := range errs {
		assert.ErrorIs(t, err, ErrParallelTest)
	}
}
