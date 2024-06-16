package app

import (
	"fmt"
	"os"
)

func fsEnsureDirsExist(paths ...string) error {
	for _, path := range paths {
		if err := os.MkdirAll(path, 0o077); err != nil {
			return fmt.Errorf("file sys: ensure dirs exist: %s", path)
		}
	}

	return nil
}
