package gotils_test

import (
	"fmt"
	"testing"

	"github.com/mreck/gotils"
	"github.com/stretchr/testify/assert"
)

func Test_Try(t *testing.T) {
	err := gotils.Try(3, 0, func(attempt int) error {
		return fmt.Errorf("test")
	})
	assert.Error(t, err)

	err = gotils.Try(3, 0, func(attempt int) error {
		if attempt < 2 {
			return fmt.Errorf("test")
		}
		return nil
	})
	assert.NoError(t, err)
}
