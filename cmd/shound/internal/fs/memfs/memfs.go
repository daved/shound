package memfs

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/liamg/memoryfs"
)

type MemFS struct {
	*memoryfs.FS
}

func New() *MemFS {
	return &MemFS{
		FS: memoryfs.New(),
	}
}

func (s *MemFS) Stat(name string) (fs.FileInfo, error) {
	fi, err := s.FS.Stat(name)
	if err != nil {
		return nil, fmt.Errorf("memfs: stat: %w", err)
	}

	return fi, nil
}

func (s *MemFS) Open(name string) (fs.File, error) {
	f, err := s.FS.Open(name)
	if err != nil {
		return nil, fmt.Errorf("memfs: open: %w", err)
	}

	return f, nil
}

func (s *MemFS) ReadFile(path string) ([]byte, error) {
	bs, err := s.FS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("memfs: readfile: %w", err)
	}

	return bs, nil
}

func (s *MemFS) WriteFile(path string, data []byte, perm os.FileMode) error {
	if err := s.FS.WriteFile(path, data, perm); err != nil {
		return fmt.Errorf("memfs: writefile: %w", err)
	}

	return nil
}

func (s *MemFS) MkdirAll(path string, perm os.FileMode) error {
	if err := s.FS.MkdirAll(path, perm); err != nil {
		return fmt.Errorf("memfs: mkdirall: %w", err)
	}

	return nil
}

func (s *MemFS) RemoveAll(path string) error {
	if err := s.FS.RemoveAll(path); err != nil {
		return fmt.Errorf("memfs: removeall: %w", err)
	}

	return nil
}
