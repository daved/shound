package main

import (
	"io"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type CmdsSounds map[string]string // map[CmdName]SoundFile

type Config struct {
	flags *GlobalFlags
	file  *ConfigFile

	help       bool
	soundDir   FilePath
	playCmd    string
	cmdsSounds CmdsSounds
}

func NewConfig() *Config {
	return &Config{
		flags: &GlobalFlags{},
		file:  &ConfigFile{},
	}
}

func (c *Config) Resolve() error {
	c.playCmd = c.file.PlayCmd
	// TODO: handle sounddir construction appropriately
	c.soundDir = FilePath(filepath.Join(string(c.file.SoundCache), string(c.file.Theme)))
	c.cmdsSounds = c.file.CmdSounds
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
