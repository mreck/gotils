package gotils

import (
	"context"
	"time"
)

// Try reruns the function until it succeeds or runs out of attempts
func Try(ctx context.Context, maxAttempts int, waitDuration time.Duration, f func(ctx context.Context, attempt int) error) error {
	var err error

	for attempt := range maxAttempts {
		err = ctx.Err()
		if err != nil {
			break
		}

		err = f(ctx, attempt)
		if err == nil {
			break
		}

		time.Sleep(waitDuration)
	}

	return err
}

// Try2 reruns the function until it succeeds or runs out of attempts
func Try2[T any](ctx context.Context, maxAttempts int, waitDuration time.Duration, f func(ctx context.Context, attempt int) (T, error)) (T, error) {
	var (
		val T
		err error
	)

	for attempt := range maxAttempts {
		err = ctx.Err()
		if err != nil {
			break
		}

		val, err = f(ctx, attempt)
		if err == nil {
			break
		}

		time.Sleep(waitDuration)
	}

	return val, err
}
