package gen

import (
	"os"
)

var setup = func(name string) (File, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return &RealFile{file}, nil
}

// File represents an interface for file operations.
type File interface {
	WriteString(s string) (n int, err error)
	Close() error
}

type RealFile struct {
	*os.File
}

// MockFile is a mock implementation of the File interface for testing.
type MockFile struct {
	WriteStringFunc func(s string) (n int, err error)
	CloseFunc func() error
}

// Write mocks the Write method.
func (m *MockFile) WriteString(s string) (n int, err error) {
	return m.WriteStringFunc(s)
}

// Close mocks the Close method.
func (m *MockFile) Close() error {
	return m.CloseFunc()
}

func Create(name string) (File, error) {
	return setup(name)
}
