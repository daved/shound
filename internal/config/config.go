package config

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/daved/shound/internal/fpath"
)

const (
	notFoundKey = "__notfound"
)

type CmdsSounds map[string]string // map[CmdName]SoundFile

type Config struct {
	UserFlags *Flags
	UserFile  *File

	*Flags

	*File

	NotFoundKey   string
	NotFoundSound string
}

func NewConfig(defConfPath string) *Config {
	return &Config{
		UserFlags:   &Flags{ConfFilePath: defConfPath},
		UserFile:    new(File),
		Flags:       new(Flags),
		File:        new(File),
		NotFoundKey: notFoundKey,
	}
}

func (c *Config) Resolve() error { // NOTE: A
	*c.Flags = *c.UserFlags
	*c.File = *c.UserFile

	c.CmdSounds = cloneMap(c.UserFile.CmdSounds)
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
	SoundCache fpath.FilePath
	Theme      string
	PlayCmd    string
	CmdSounds  map[string]string
}

func (c *File) InitFromTOML(r io.Reader) error { // TODO: handle errors | A
	if _, err := toml.NewDecoder(r).Decode(&c); err != nil {
		return err
	}

	if err := c.SoundCache.Validate(); err != nil {
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
