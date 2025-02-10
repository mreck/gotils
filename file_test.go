package gotils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mreck/gotils"
)

func TestFilePathExists(t *testing.T) {
	var exists bool
	var err error

	exists, err = gotils.FilePathExists("file.go")
	assert.True(t, exists)
	assert.Nil(t, err)

	exists, err = gotils.FilePathExists("file_123.go")
	assert.False(t, exists)
	assert.Nil(t, err)
}
