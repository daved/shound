package config

import (
	"fmt"
	"path/filepath"

	"github.com/daved/shound/internal/config"
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
	User     *User
	Resolved *config.Config
}

func New(defConfPath, defThemesDirPath string) *Config {
	return &Config{
		User: &User{
			Flags: &Flags{
				ConfFilePath: defConfPath,
			},
			File: &File{
				ThemesDir: defThemesDirPath,
			},
			ThemeFile: new(ThemeFile),
		},
		Resolved: &config.Config{},
	}
}

func (c *Config) ValidateFiles() error {
	eMsg := "config: validate: %w"

	if err := c.User.File.validate(); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (c *Config) Resolve() error {
	c.Resolved.Help = c.User.Flags.Help
	c.Resolved.ConfFilePath = c.User.Flags.ConfFilePath
	c.Resolved.Bypass = c.User.File.Bypass
	c.Resolved.PlayCmd = c.User.File.PlayCmd
	c.Resolved.ThemesDir = c.User.File.ThemesDir
	c.Resolved.ThemeDir = filepath.Join(string(c.User.File.ThemesDir), string(c.User.File.ThemeRepo))
	c.Resolved.ThemeRepo = c.User.File.ThemeRepo
	c.Resolved.CmdSounds = cloneMap(c.User.ThemeFile.CmdSounds)
	c.Resolved.NotFoundKey = notFoundKey

	overrides, ok := c.User.File.ThemeOverrides[c.User.File.ThemeRepo]
	if ok {
		for k, v := range overrides {
			c.Resolved.CmdSounds[k] = v
		}
	}

	if v, ok := c.Resolved.CmdSounds[notFoundKey]; ok {
		c.Resolved.NotFoundSound = v
		delete(c.Resolved.CmdSounds, notFoundKey)
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
