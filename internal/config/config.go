package config

import (
	"io"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/daved/shound/internal/fpath"
)

const (
	notFoundKey = "__notfound"
)

type CmdsSounds map[string]string // map[CmdName]SoundFile

type Config struct {
	Flags *GlobalFlags
	File  *ConfigFile

	Help       bool
	SoundDir   fpath.FilePath
	PlayCmd    string
	CmdsSounds CmdsSounds
	NoCmdSound string
}

func NewConfig() *Config {
	return &Config{
		Flags: &GlobalFlags{},
		File:  &ConfigFile{},
	}
}

func (c *Config) Resolve() error { // NOTE: A
	c.Help = c.Flags.Help
	c.PlayCmd = c.File.PlayCmd

	c.CmdsSounds = cloneMap(c.File.CmdSounds)
	for k, v := range c.CmdsSounds {
		if k == notFoundKey {
			c.NoCmdSound = v
			delete(c.CmdsSounds, k)
		}
	}

	// TODO: A: handle sounddir construction appropriately
	c.SoundDir = fpath.FilePath(filepath.Join(string(c.File.SoundCache), string(c.File.Theme)))

	return nil
}

type GlobalFlags struct {
	Help bool
}

type ConfigFile struct {
	SoundCache fpath.FilePath
	Theme      string
	PlayCmd    string
	CmdSounds  map[string]string
}

func (c *ConfigFile) InitFromTOML(r io.Reader) error { // TODO: handle errors | A
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
