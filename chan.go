package gotils

import (
	"context"
)

// CollectChannelMessages pulls all messages from the channel, until the channel is closed or the context runs out
func CollectChannelMessages[T any](ctx context.Context, c chan T) ([]T, error) {
	var msgs []T
	var stop bool

	for {
		select {
		case m, ok := <-c:
			if !ok {
				stop = true
				break
			}
			msgs = append(msgs, m)
		case <-ctx.Done():
			return msgs, ctx.Err()
		}

		if stop {
			break
		}
	}

	return msgs, nil
}
