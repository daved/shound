package config

import (
	"errors"

	"github.com/daved/shound/internal/fpath"
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
	UserFlags *Flags
	UserFile  *File

	*Flags

	// File
	Active        bool
	PlayCmd       string
	ThemesDir     fpath.FilePath
	ThemeName     string
	ThemeFallback string
	CmdSounds     CmdSounds

	NotFoundKey   string
	NotFoundSound string
}

func NewConfig(defConfPath string) *Config {
	return &Config{
		UserFlags:   &Flags{ConfFilePath: defConfPath},
		UserFile:    new(File),
		Flags:       new(Flags),
		NotFoundKey: notFoundKey,
	}
}

func (c *Config) Resolve() error { // NOTE: A
	*c.Flags = *c.UserFlags

	c.Active = c.UserFile.Active
	c.PlayCmd = c.UserFile.PlayCmd
	c.ThemesDir = c.UserFile.ThemesDir
	c.ThemeName = c.UserFile.ThemeName
	c.ThemeFallback = c.UserFile.ThemeFallback

	cmdSounds, ok := c.UserFile.ThemeOverrides[c.ThemeName]
	if !ok {
		return errors.New("theme not found")
	}
	c.CmdSounds = cloneMap(cmdSounds)

	for k, v := range c.CmdSounds {
		if k == c.NotFoundKey {
			c.NotFoundSound = v
			delete(c.CmdSounds, k)
		}

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
	ThemesDir      fpath.FilePath `yaml:"ThemesDir"`
	ThemeName      string         `yaml:"ThemeName"`
	ThemeFallback  string         `yaml:"ThemeFallback"`
	ThemeOverrides ThemeOverrides `yaml:"CmdSoundsOverrides"`
}

func (c *File) InitFromYAML(data []byte) error { // TODO: handle errors | A
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}

	if err := c.ThemesDir.Validate(); err != nil {
		return err
	}

	// TODO: A: validate combination of soundcache+theme

	return nil
}

func cloneMap(in map[string]string) map[string]string {
	out := make(map[string]string)

	for k, v := range in {
		out[k] = v
	}

	return out
}
