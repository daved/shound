package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/daved/shound/internal/config"
)

type themesInfo struct {
	cnf           *config.Config
	themesDir     string
	themeFileName string
}

func newThemesInfo(cnf *config.Config, fileName string) *themesInfo {
	return &themesInfo{
		cnf:           cnf,
		themesDir:     cnf.User.File.ThemesDir,
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
	eMsg := "themes info: set theme: %w"

	cnfBytes, err := os.ReadFile(i.cnf.User.Flags.ConfFilePath)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	updCnfBytes, err := config.SetFileField(cnfBytes, config.FileFieldThemeRepo, theme)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := os.WriteFile(i.cnf.User.Flags.ConfFilePath, updCnfBytes, 0600); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := i.cnf.User.File.InitFromYAML(cnfBytes); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
