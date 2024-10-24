package fs

import (
	"io/fs"
	"os"
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
