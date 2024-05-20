package main

import (
	"io"

	"github.com/BurntSushi/toml"
)

type Config struct {
	flags *GlobalFlags
	file  *ConfigFile

	help       bool
	audioCache FilePath
	playCmd    string
	cmdSounds  map[string]string
}

func NewConfig() *Config {
	return &Config{
		flags: &GlobalFlags{},
		file:  &ConfigFile{},
	}
}

func (c *Config) Resolve() error {
	// TODO: fill
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
