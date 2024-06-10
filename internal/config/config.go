package config

import (
	"path/filepath"
)

const (
	notFoundKey = "__notfound"
)

type Flags struct {
	Help         bool
	ConfFilePath string
}

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
	ThemesDir string
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
	c.ThemesDir = c.User.File.ThemesDir
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

func cloneMap(in map[string]string) map[string]string {
	out := make(map[string]string)

	for k, v := range in {
		out[k] = v
	}

	return out
}
