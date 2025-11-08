package gotils

import (
	"context"
	"sync"
)

// ParellelFor iterates over the slice in parallel using Goroutines.
func ParellelFor[T any](ctx context.Context, values []T, coroutines int, fn func(context.Context, int, T) error) []error {
	type dataT struct {
		idx int
		val T
	}

	var (
		wg    sync.WaitGroup
		m     sync.Mutex
		errs  []error
		dataC = make(chan dataT)
	)

	for range coroutines {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					m.Lock()
					errs = append(errs, ctx.Err())
					m.Unlock()
					return

				case msg, ok := <-dataC:
					if !ok {
						return
					}
					err := fn(ctx, msg.idx, msg.val)
					if err != nil {
						m.Lock()
						errs = append(errs, err)
						m.Unlock()
					}
				}
			}
		}()
	}

	for i, d := range values {
		// @TODO: Check if the context is done.
		dataC <- dataT{i, d}
	}

	close(dataC)
	wg.Wait()

	return errs
}

// ParellelMap maps over the slice in parallel using Goroutines.
func ParellelMap[T any, R any](ctx context.Context, values []T, coroutines int, fn func(context.Context, int, T) (R, error)) ([]R, []error) {
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

	var wgData sync.WaitGroup

	for range coroutines {
		wgData.Add(1)

		go func() {
			defer wgData.Done()

			for {
				select {
				case <-ctx.Done():
					// @TODO: handle properly
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
	var wgRes sync.WaitGroup

	wgRes.Add(1)
	go func() {
		defer wgRes.Done()

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
	wgData.Wait()

	close(resC)
	wgRes.Wait()

	return result, errs
}
