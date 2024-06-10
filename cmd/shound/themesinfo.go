package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type themesInfo struct {
	themesDir     string
	themeFileName string
}

func newThemesInfo(dir, fileName string) *themesInfo {
	return &themesInfo{
		themesDir:     dir,
		themeFileName: fileName,
	}
}

func (i *themesInfo) Themes() ([]string, error) {
	var ts []string

	err := filepath.WalkDir(i.themesDir, func(path string, de fs.DirEntry, err error) error {
		if filepath.Base(path) == i.themeFileName {
			dir := filepath.Dir(path)
			relToThemesdir := dir[len(i.themesDir):]
			noSeps := strings.Trim(relToThemesdir, string(os.PathSeparator))
			ts = append(ts, noSeps)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("themes info: themes: %w", err)
	}

	return ts, nil
}

func (i *themesInfo) SetTheme(theme string) error {
	return nil
}
