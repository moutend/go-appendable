// Package appendable provides File struct supports append operation.
package appendable

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// File is a file supports append operation.
type File struct {
	name string
	*bytes.Buffer
}

// Close performs os.File.Close() method on underlying file.
func (f *File) Close() error {
	file, err := os.Create(f.name)

	if err != nil {
		return fmt.Errorf("appendable: failed to create: %w", err)
	}

	defer file.Close()

	if _, err := io.Copy(file, f.Buffer); err != nil {
		return fmt.Errorf("appendable: failed to write: %w", err)
	}

	return nil
}

// NewFile returns appendable file.
func NewFile(name string) (*File, error) {
	original, err := os.Open(name)

	if err != nil {
		if os.IsNotExist(err) {
			return &File{name, &bytes.Buffer{}}, nil
		}

		return nil, fmt.Errorf("appendable: failed to open: %w", err)
	}

	defer original.Close()

	buffer := &bytes.Buffer{}

	if _, err := io.Copy(buffer, original); err != nil {
		return nil, fmt.Errorf("appendable: failed to read existing file: %w", err)
	}

	return &File{name, buffer}, nil
}
