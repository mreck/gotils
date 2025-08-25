package gotils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClamp(t *testing.T) {
	assert.Equal(t, 0, Clamp(0, 9, -1))
	assert.Equal(t, 9, Clamp(0, 9, 10))
	assert.Equal(t, 5, Clamp(0, 9, 5))
}
