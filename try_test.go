package gotils_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/mreck/gotils"

	"github.com/stretchr/testify/assert"
)

func Test_Try(t *testing.T) {
	ctx := context.Background()

	err := gotils.Try(ctx, 3, 0, func(ctx context.Context, attempt int) error {
		return fmt.Errorf("test")
	})

	assert.Error(t, err)

	err = gotils.Try(ctx, 3, 0, func(ctx context.Context, attempt int) error {
		if attempt < 2 {
			return fmt.Errorf("test")
		}
		return nil
	})

	assert.NoError(t, err)

	ctx, cancel := context.WithCancel(ctx)
	cancel()

	err = gotils.Try(ctx, 3, 0, func(ctx context.Context, attempt int) error {
		return nil
	})

	assert.ErrorIs(t, err, context.Canceled)
}

func Test_Try2(t *testing.T) {
	ctx := context.Background()

	n, err := gotils.Try2(ctx, 3, 0, func(ctx context.Context, attempt int) (int, error) {
		return attempt, fmt.Errorf("test")
	})

	assert.Error(t, err)
	assert.Equal(t, 2, n)

	n, err = gotils.Try2(ctx, 3, 0, func(ctx context.Context, attempt int) (int, error) {
		if attempt < 1 {
			return attempt, fmt.Errorf("test")
		}
		return attempt, nil
	})

	assert.NoError(t, err)
	assert.Equal(t, 1, n)

	ctx, cancel := context.WithCancel(ctx)
	cancel()

	n, err = gotils.Try2(ctx, 3, 0, func(ctx context.Context, attempt int) (int, error) {
		if attempt < 1 {
			return attempt, fmt.Errorf("test")
		}
		return attempt, nil
	})

	assert.ErrorIs(t, err, context.Canceled)
	assert.Equal(t, 0, n)
}
