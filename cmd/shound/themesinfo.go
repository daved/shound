package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/daved/shound/internal/config"
	"github.com/go-git/go-git/v5"
)

type themesInfo struct {
	out           io.Writer
	cnf           *config.Config
	themesDir     string
	themeFileName string
}

func newThemesInfo(out io.Writer, cnf *config.Config, fileName string) *themesInfo {
	return &themesInfo{
		out:           out,
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

func (i *themesInfo) AddTheme(theme string) error {
	eMsg := "themes info: add theme: %w"

	themePath := filepath.Join(i.themesDir, theme)

	if err := os.MkdirAll(themePath, 0755); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	// TODO: check if already installed
	// TODO: if version different, checkout correct version

	// TODO: handle versions on initial clone
	// TODO: load to tmp dir and validate before moving to themes dir
	_, err := git.PlainClone(themePath, false, &git.CloneOptions{
		URL:      fmt.Sprintf("https://%s", theme),
		Progress: i.out,
	})
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (i *themesInfo) DeleteTheme(theme string) error {
	eMsg := "themes info: delete theme: %w"

	themePath := filepath.Join(i.themesDir, theme)
	if err := os.RemoveAll(themePath); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
