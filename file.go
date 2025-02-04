package gotils

import (
	"errors"
	"os"
)

// FilePathExists check if a file (or directory) exists
func FilePathExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
