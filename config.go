package main

import (
	"io"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	notFoundKey = "__notfound"
)

type CmdsSounds map[string]string // map[CmdName]SoundFile

type Config struct {
	flags *GlobalFlags
	file  *ConfigFile

	help       bool
	soundDir   FilePath
	playCmd    string
	cmdsSounds CmdsSounds
	noCmdSound string
}

func NewConfig() *Config {
	return &Config{
		flags: &GlobalFlags{},
		file:  &ConfigFile{},
	}
}

func (c *Config) Resolve() error {
	c.playCmd = c.file.PlayCmd

	c.cmdsSounds = cloneMap(c.file.CmdSounds)
	for k, v := range c.cmdsSounds {
		if k == notFoundKey {
			c.noCmdSound = v
			delete(c.cmdsSounds, k)
		}
	}

	// TODO: handle sounddir construction appropriately
	c.soundDir = FilePath(filepath.Join(string(c.file.SoundCache), string(c.file.Theme)))

	return nil
}

type GlobalFlags struct {
	help bool
}

type ConfigFile struct {
	SoundCache FilePath
	Theme      string
	PlayCmd    string
	CmdSounds  map[string]string
}

func (c *ConfigFile) initFromTOML(r io.Reader) error {
	if _, err := toml.NewDecoder(r).Decode(&c); err != nil {
		// TODO: error
		return err
	}

	if err := c.SoundCache.Validate(); err != nil {
		// TODO: error
		return err
	}

	// TODO: validate combination of soundcache+theme

	return nil
}

func cloneMap(in map[string]string) map[string]string {
	out := make(map[string]string)

	for k, v := range in {
		out[k] = v
	}

	return out
}
