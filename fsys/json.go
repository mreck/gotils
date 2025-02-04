package fs

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

// ReadJSONFile reads and decodes the file content into the destination
func ReadJSONFile(filename string, dst any) error {
	fp, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("opening file failed: %w", err)
	}
	defer fp.Close()

	err = json.NewDecoder(fp).Decode(dst)
	if err != nil {
		return fmt.Errorf("decoding json failed: %w", err)
	}

	return nil
}

func WriteJSONFile(filename string, indent string, data any, perm ...fs.FileMode) error {
	actualPerm := os.ModePerm
	if len(perm) > 0 {
		actualPerm = perm[0]
	}

	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, actualPerm)
	if err != nil {
		return fmt.Errorf("opening file failed: %w", err)
	}
	defer fp.Close()

	enc := json.NewEncoder(fp)
	if len(indent) > 0 {
		enc.SetIndent("", indent)
	}

	err = enc.Encode(data)
	if err != nil {
		return fmt.Errorf("encoding data failed: %w", err)
	}

	return nil
}
