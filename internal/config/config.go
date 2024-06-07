package config

import (
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	notFoundKey = "__notfound"
)

type CmdSounds map[string]string // map[CommandName]SoundFile

func (css CmdSounds) CmdList() []string {
	cmds := make([]string, 0, len(css))
	for cmd := range css {
		cmds = append(cmds, cmd)
	}
	return cmds
}

type ThemeOverrides map[string]map[string]string // map[ThemeRepo]map[CommandName]SoundFile

type User struct {
	Flags     *Flags
	File      *File
	ThemeFile *ThemeFile
}

type Config struct {
	User *User

	*Flags

	// from file
	Active    bool
	PlayCmd   string
	ThemeDir  string
	ThemeRepo string

	// from theme file
	CmdSounds     CmdSounds
	NotFoundKey   string
	NotFoundSound string
}

func NewConfig(defConfPath string) *Config {
	return &Config{
		User: &User{
			Flags:     &Flags{ConfFilePath: defConfPath},
			File:      new(File),
			ThemeFile: new(ThemeFile),
		},
		Flags:       new(Flags),
		NotFoundKey: notFoundKey,
	}
}

func (c *Config) Resolve() error {
	*c.Flags = *c.User.Flags

	c.Active = c.User.File.Active
	c.PlayCmd = c.User.File.PlayCmd
	c.ThemeDir = filepath.Join(string(c.User.File.ThemesDir), string(c.User.File.ThemeRepo))
	c.ThemeRepo = c.User.File.ThemeRepo

	c.CmdSounds = cloneMap(c.User.ThemeFile.CmdSounds)
	overrides, ok := c.User.File.ThemeOverrides[c.User.File.ThemeRepo]
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
