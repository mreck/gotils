package gotils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

var (
	ErrDecodingFile = errors.New("decoding file failed")
	ErrEncodingFile = errors.New("encoding file failed")
	ErrOpeningFile  = errors.New("opening file failed")
)

// ReadJSONFile reads and decodes the file content into the destination
func ReadJSONFile(filename string, dst any) error {
	fp, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrOpeningFile, err)
	}
	defer fp.Close()

	err = json.NewDecoder(fp).Decode(dst)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrDecodingFile, err)
	}

	return nil
}

func WriteJSONFile(filename string, indent string, data any, perm ...fs.FileMode) error {
	actualPerm := os.ModePerm
	if len(perm) > 0 {
		actualPerm = perm[0]
	}

	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, actualPerm)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrOpeningFile, err)
	}
	defer fp.Close()

	enc := json.NewEncoder(fp)
	if len(indent) > 0 {
		enc.SetIndent("", indent)
	}

	err = enc.Encode(data)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrEncodingFile, err)
	}

	return nil
}
