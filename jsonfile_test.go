package gotils

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJSONFile(t *testing.T) {
	var dInt []int
	var dStr []string
	var err error

	err = ReadJSONFile("test.json", &dInt)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, dInt)

	err = ReadJSONFile("test_123.json", &dInt)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrOpeningFile))

	err = ReadJSONFile("test.json", &dStr)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrDecodingFile))
}

func TestWriteJSONFile(t *testing.T) {
	var b []byte
	var err error

	fn := "tmp.json"

	err = WriteJSONFile(fn, "", []int{1, 2, 3})
	assert.Nil(t, err)
	b, err = os.ReadFile(fn)
	assert.Nil(t, err)
	assert.Equal(t, []byte("[1,2,3]\n"), b)

	err = WriteJSONFile(fn, "\t", []int{1, 2, 3})
	assert.Nil(t, err)
	b, err = os.ReadFile(fn)
	assert.Nil(t, err)
	assert.Equal(t, []byte("[\n\t1,\n\t2,\n\t3\n]\n"), b)
}
