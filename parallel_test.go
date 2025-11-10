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

func Test_ParallelFor(t *testing.T) {
	var (
		out []int
		m   sync.Mutex
	)

	canceled, cancel := context.WithCancel(context.Background())
	cancel()

	testCases := []struct {
		ctx          context.Context
		fn           func(ctx context.Context, i int, n int) error
		expectedOut  []int
		expectedErrs []error
	}{
		{
			ctx: context.Background(),
			fn: func(ctx context.Context, i int, n int) error {
				m.Lock()
				defer m.Unlock()
				out = append(out, n+1)
				return nil
			},
			expectedOut:  []int{2, 3, 4, 5, 6, 7},
			expectedErrs: nil,
		},
		{
			ctx: context.Background(),
			fn: func(ctx context.Context, i int, n int) error {
				if n%2 == 0 {
					return ErrParallelTest
				}
				m.Lock()
				defer m.Unlock()
				out = append(out, n+1)
				return nil
			},
			expectedOut: []int{2, 4, 6},
			expectedErrs: []error{
				ErrParallelTest,
				ErrParallelTest,
				ErrParallelTest,
			},
		},
		{
			ctx: canceled,
			fn: func(ctx context.Context, i int, n int) error {
				m.Lock()
				defer m.Unlock()
				out = append(out, n+1)
				return nil
			},
			expectedOut: []int(nil),
			expectedErrs: []error{
				context.Canceled,
				context.Canceled,
				context.Canceled,
				context.Canceled,
				context.Canceled,
				context.Canceled,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			in := []int{1, 2, 3, 4, 5, 6}
			out = []int(nil)

			errs := gotils.ParallelFor(tc.ctx, in, 4, tc.fn)

			sort.Ints(out)
			assert.Equal(t, tc.expectedErrs, errs)
			assert.Equal(t, tc.expectedOut, out)
		})
	}
}

func Test_ParallelMap(t *testing.T) {
	canceled, cancel := context.WithCancel(context.Background())
	cancel()

	testCases := []struct {
		ctx      context.Context
		fn       func(ctx context.Context, i int, n int) (int, error)
		expected []gotils.Result[int]
	}{
		{
			ctx: context.Background(),
			fn: func(ctx context.Context, i int, n int) (int, error) {
				return n + 1, nil
			},
			expected: []gotils.Result[int]{
				{2, nil},
				{3, nil},
				{4, nil},
				{5, nil},
				{6, nil},
				{7, nil},
			},
		},
		{
			ctx: context.Background(),
			fn: func(ctx context.Context, i int, n int) (int, error) {
				if n%2 == 0 {
					return 0, fmt.Errorf("%w: %d/%d", ErrParallelTest, i, n)
				}
				return n, nil
			},
			expected: []gotils.Result[int]{
				{1, nil},
				{0, fmt.Errorf("%w: %d/%d", ErrParallelTest, 1, 2)},
				{3, nil},
				{0, fmt.Errorf("%w: %d/%d", ErrParallelTest, 3, 4)},
				{5, nil},
				{0, fmt.Errorf("%w: %d/%d", ErrParallelTest, 5, 6)},
			},
		},
		{
			ctx: canceled,
			fn: func(ctx context.Context, i int, n int) (int, error) {
				return n + 1, nil
			},
			expected: []gotils.Result[int]{
				{0, context.Canceled},
				{0, context.Canceled},
				{0, context.Canceled},
				{0, context.Canceled},
				{0, context.Canceled},
				{0, context.Canceled},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			in := []int{1, 2, 3, 4, 5, 6}
			res := gotils.ParallelMap(tc.ctx, in, 4, tc.fn)

			assert.Equal(t, tc.expected, res)
		})
	}
}
