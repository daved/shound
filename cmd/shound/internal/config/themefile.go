package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type CmdSounds map[string]string // map[CommandName]SoundFile

func (css CmdSounds) CmdList() []string {
	cmds := make([]string, 0, len(css))
	for cmd := range css {
		cmds = append(cmds, cmd)
	}
	return cmds
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
