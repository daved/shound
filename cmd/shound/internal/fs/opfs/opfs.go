package opfs

import (
	"fmt"
	"io/fs"
	"os"
)

type OpFS struct{}

func New() *OpFS {
	return &OpFS{}
}

func (s *OpFS) Stat(name string) (fs.FileInfo, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return nil, fmt.Errorf("opfs: stat: %w", err)
	}

	return fi, nil
}

func (s *OpFS) Open(name string) (fs.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("opfs: open: %w", err)
	}

	return f, nil
}

func (s *OpFS) ReadFile(path string) ([]byte, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("opfs: readfile: %w", err)
	}

	return bs, nil
}

func (s *OpFS) WriteFile(path string, data []byte, perm os.FileMode) error {
	if err := os.WriteFile(path, data, perm); err != nil {
		return fmt.Errorf("opfs: writefile: %w", err)
	}

	return nil
}

func (s *OpFS) MkdirAll(path string, perm os.FileMode) error {
	if err := os.MkdirAll(path, perm); err != nil {
		return fmt.Errorf("opfs: mkdirall: %w", err)
	}

	return nil
}

func (s *OpFS) RemoveAll(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("opfs: removeall: %w", err)
	}

	return nil
}
