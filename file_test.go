package gotils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilePathExists(t *testing.T) {
	var exists bool
	var err error

	exists, err = FilePathExists("file.go")
	assert.True(t, exists)
	assert.Nil(t, err)

	exists, err = FilePathExists("file_123.go")
	assert.False(t, exists)
	assert.Nil(t, err)
}
