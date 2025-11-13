package gotils

import (
	"context"
	"sync"
)

// ParallelFor iterates over the slice in parallel using Goroutines.
func ParallelFor[T any](ctx context.Context, values []T, coroutines int, fn func(context.Context, int, T) error) []error {
	type dataT struct {
		idx int
		val T
	}

	var (
		m     sync.Mutex
		wg    sync.WaitGroup
		errs  []error
		dataC = make(chan dataT)
	)

	for range coroutines {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				msg, ok := <-dataC
				if !ok {
					return
				}

				err := ctx.Err()
				if err != nil {
					m.Lock()
					errs = append(errs, ctx.Err())
					m.Unlock()
					continue
				}

				err = fn(ctx, msg.idx, msg.val)
				if err != nil {
					m.Lock()
					errs = append(errs, err)
					m.Unlock()
				}
			}
		}()
	}

	for i, d := range values {
		dataC <- dataT{i, d}
	}

	close(dataC)
	wg.Wait()

	return errs
}

// ParallelMap maps over the slice in parallel using Goroutines.
func ParallelMap[T any, R any](ctx context.Context, values []T, coroutines int, fn func(context.Context, int, T) (R, error)) []Result[R] {
	type dataT struct {
		idx int
		val T
	}

	var (
		m      sync.Mutex
		wg     sync.WaitGroup
		result = make([]Result[R], len(values))
		dataC  = make(chan dataT)
	)

	for range coroutines {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				msg, ok := <-dataC
				if !ok {
					return
				}

				err := ctx.Err()
				if err != nil {
					m.Lock()
					result[msg.idx].Err = err
					m.Unlock()
					continue
				}

				val, err := fn(ctx, msg.idx, msg.val)
				m.Lock()
				result[msg.idx].Value = val
				result[msg.idx].Err = err
				m.Unlock()
			}
		}()
	}

	for i, d := range values {
		dataC <- dataT{i, d}
	}

	close(dataC)
	wg.Wait()

	return result
}

// @TODO: ParallelReduce
