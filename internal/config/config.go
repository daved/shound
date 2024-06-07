package config

import (
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	notFoundKey = "__notfound"
)

type (
	CmdSounds      map[string]string            // map[CommandName]SoundFile
	ThemeOverrides map[string]map[string]string // map[ThemeName]map[CommandName]SoundFile
)

type Config struct {
	// untouched input values (restrict access to setup)

	UserFlags *Flags

	UserFile *File

	ThemeFile *ThemeFile

	// resolved values (use after resolve method is called)

	*Flags

	// from file
	Active   bool
	PlayCmd  string
	ThemeDir string

	// from theme file
	CmdSounds     CmdSounds
	NotFoundKey   string
	NotFoundSound string
}

func NewConfig(defConfPath string) *Config {
	return &Config{
		UserFlags:   &Flags{ConfFilePath: defConfPath},
		UserFile:    new(File),
		ThemeFile:   new(ThemeFile),
		Flags:       new(Flags),
		NotFoundKey: notFoundKey,
	}
}

func (c *Config) Resolve() error {
	*c.Flags = *c.UserFlags

	c.Active = c.UserFile.Active
	c.PlayCmd = c.UserFile.PlayCmd
	c.ThemeDir = filepath.Join(string(c.UserFile.ThemesDir), string(c.UserFile.ThemeName))

	c.CmdSounds = cloneMap(c.ThemeFile.CmdSounds)
	overrides, ok := c.UserFile.ThemeOverrides[c.UserFile.ThemeName]
	if ok {
		for k, v := range overrides {
			c.CmdSounds[k] = v
		}
	}

	if v, ok := c.CmdSounds[c.NotFoundKey]; ok {
		c.NotFoundSound = v
		delete(c.CmdSounds, c.NotFoundKey)
	}

	return nil
}

type Flags struct {
	Help         bool
	ConfFilePath string
}

type File struct {
	Active         bool           `yaml:"Active"`
	PlayCmd        string         `yaml:"PlayCmd"`
	ThemesDir      string         `yaml:"ThemesDir"`
	ThemeName      string         `yaml:"ThemeName"`
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
	return filepath.Join(string(f.ThemesDir), f.ThemeName, fileName)
}

type ThemeFile struct {
	CmdSounds CmdSounds `yaml:"CmdSounds"`
}

func (f *ThemeFile) InitFromYAML(data []byte) error {
	eMsg := "config: theme file: init from yaml: %w"

	if err := yaml.Unmarshal(data, f); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func cloneMap(in map[string]string) map[string]string {
	out := make(map[string]string)

	for k, v := range in {
		out[k] = v
	}

	return out
}
