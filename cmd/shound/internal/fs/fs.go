package fs

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/liamg/memoryfs"
)

var (
	_ FS = (*OpFS)(nil)
	_ FS = (*MemFS)(nil)
)

type (
	DirEntry = fs.DirEntry
)

type FS interface {
	fs.StatFS
	fs.ReadFileFS
	MkdirAll(path string, perm os.FileMode) error
	WriteFile(path string, data []byte, perm os.FileMode) error
	RemoveAll(path string) error
}

type OpFS struct {
}

func NewOpFS() *OpFS {
	return &OpFS{}
}

func (s *OpFS) Stat(name string) (fs.FileInfo, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return nil, fmt.Errorf("fs: opfs: stat: %w", err)
	}

	return fi, nil
}

func (s *OpFS) Open(name string) (fs.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("fs: opfs: open: %w", err)
	}

	return f, nil
}

func (s *OpFS) ReadFile(path string) ([]byte, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("fs: opfs: %w", err)
	}

	return bs, nil
}

func (s *OpFS) WriteFile(path string, data []byte, perm os.FileMode) error {
	if err := os.WriteFile(path, data, perm); err != nil {
		return fmt.Errorf("fs: opfs: writefile: %w", err)
	}

	return nil
}

func (s *OpFS) MkdirAll(path string, perm os.FileMode) error {
	if err := os.MkdirAll(path, perm); err != nil {
		return fmt.Errorf("fs: opfs: mkdirall: %w", err)
	}

	return nil
}

func (s *OpFS) RemoveAll(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("fs: opfs: removeall: %w", err)
	}

	return nil
}

type MemFS struct {
	*memoryfs.FS
}

func NewMemFS() *MemFS {
	return &MemFS{
		FS: memoryfs.New(),
	}
}

func (s *MemFS) Stat(name string) (fs.FileInfo, error) {
	fi, err := s.FS.Stat(name)
	if err != nil {
		return nil, fmt.Errorf("fs: memfs: stat: %w", err)
	}

	return fi, nil
}

func (s *MemFS) Open(name string) (fs.File, error) {
	f, err := s.FS.Open(name)
	if err != nil {
		return nil, fmt.Errorf("fs: memfs: open: %w", err)
	}

	return f, nil
}

func (s *MemFS) ReadFile(path string) ([]byte, error) {
	bs, err := s.FS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("fs: memfs: %w", err)
	}

	return bs, nil
}

func (s *MemFS) WriteFile(path string, data []byte, perm os.FileMode) error {
	if err := s.FS.WriteFile(path, data, perm); err != nil {
		return fmt.Errorf("fs: memfs: writefile: %w", err)
	}

	return nil
}

func (s *MemFS) MkdirAll(path string, perm os.FileMode) error {
	if err := s.FS.MkdirAll(path, perm); err != nil {
		return fmt.Errorf("fs: memfs: mkdirall: %w", err)
	}

	return nil
}

func (s *MemFS) RemoveAll(path string) error {
	if err := s.FS.RemoveAll(path); err != nil {
		return fmt.Errorf("fs: memfs: removeall: %w", err)
	}

	return nil
}
