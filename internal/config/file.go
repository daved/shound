package config

import (
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ThemeOverrides map[string]map[string]string // map[ThemeRepo]map[CommandName]SoundFile

type File struct {
	Active         bool           `yaml:"Active"`
	PlayCmd        string         `yaml:"PlayCmd"`
	ThemesDir      string         `yaml:"ThemesDir"`
	ThemeRepo      string         `yaml:"ThemeRepo"`
	ThemeOverrides ThemeOverrides `yaml:"CmdSoundsOverrides"`
}

func (f *File) InitFromYAML(data []byte) error {
	eMsg := "config: file: init from yaml: %w"

	if err := yaml.Unmarshal(data, f); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (f *File) ThemePath(fileName string) string {
	return filepath.Join(string(f.ThemesDir), f.ThemeRepo, fileName)
}
