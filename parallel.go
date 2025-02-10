package gotils

import (
	"context"
	"sync"
)

// ParellelFor iterates over the slice in parallel using Goroutines.
func ParellelFor[T any](ctx context.Context, values []T, threads int, fn func(context.Context, int, T) error) []error {
	type dataT struct {
		idx int
		val T
	}

	valC := make(chan dataT)
	errC := make(chan error)

	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case msg, ok := <-valC:
					if !ok {
						return
					}
					err := fn(ctx, msg.idx, msg.val)
					if err != nil {
						errC <- err
					}
				}
			}
		}()
	}

	var errs []error
	var errWg sync.WaitGroup

	errWg.Add(1)
	go func() {
		defer errWg.Done()
		for {
			err, ok := <-errC
			if !ok {
				return
			}
			if err != nil {
				errs = append(errs, err)
			}
		}
	}()

	for i, d := range values {
		// @TODO: Check if the context is done.
		valC <- dataT{i, d}
	}

	close(valC)
	wg.Wait()
	close(errC)
	errWg.Wait()

	return errs
}

// ParellelMap maps over the slice in parallel using Goroutines.
func ParellelMap[T any, R any](ctx context.Context, values []T, threads int, fn func(context.Context, int, T) (R, error)) ([]R, []error) {
	type dataT struct {
		idx int
		val T
	}

	type resultT struct {
		idx int
		res R
		err error
	}

	valC := make(chan dataT)
	resC := make(chan resultT)

	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case msg, ok := <-valC:
					if !ok {
						return
					}
					res, err := fn(ctx, msg.idx, msg.val)
					resC <- resultT{msg.idx, res, err}
				}
			}
		}()
	}

	result := make([]R, len(values))
	var errs []error

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-resC:
				if !ok {
					return
				}
				if msg.err != nil {
					errs = append(errs, msg.err)
				}
				result[msg.idx] = msg.res
			}
		}
	}()

	for i, d := range values {
		// @TODO: Check if the context is done.
		valC <- dataT{i, d}
	}

	close(valC)
	wg.Wait()
	close(resC)

	return result, errs
}
