package app

import (
	"fmt"

	"github.com/daved/shound/cmd/shound/internal/fs"
)

func ensureDirsExist(fs fs.FS, paths ...string) error {
	for _, path := range paths {
		if err := fs.MkdirAll(path, 0o077); err != nil {
			return fmt.Errorf("file sys: ensure dirs exist: %s", path)
		}
	}

	return nil
}
