package gotils_test

import (
	"context"
	"testing"

	"github.com/mreck/gotils"

	"github.com/stretchr/testify/assert"
)

func Test_CollectChannelMessages(t *testing.T) {
	c := make(chan int, 10)
	c <- 1
	c <- 7
	c <- 3
	c <- 5
	close(c)

	ctx := context.Background()
	msgs, err := gotils.CollectChannelMessages(ctx, c)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 7, 3, 5}, msgs)
}
