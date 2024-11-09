package config

import (
	"fmt"
	"path/filepath"

	"github.com/daved/shound/internal/config"
)

const (
	notFoundKey = "__notfound"
)

type Sourced struct {
	Flags     *Flags
	File      *File
	ThemeFile *ThemeFile
	Resolved  *config.Config
	AResolved *Config
}

func NewSourced(defConfPath, defThemesDirPath string) *Sourced {
	return &Sourced{
		Flags: &Flags{
			ConfFilePath: defConfPath,
		},
		File: &File{
			ThemesDir: defThemesDirPath,
		},
		ThemeFile: new(ThemeFile),
		Resolved:  &config.Config{},
		AResolved: &Config{},
	}
}

func (s *Sourced) ValidateFiles() error {
	eMsg := "config: validate: %w"

	if err := s.File.validate(); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (s *Sourced) Resolve() error {
	s.Resolved.Help = s.Flags.Help
	s.AResolved.help = s.Flags.Help

	s.Resolved.ConfFilePath = s.Flags.ConfFilePath
	s.Resolved.Bypass = s.File.Bypass
	s.Resolved.PlayCmd = s.File.PlayCmd
	s.AResolved.playCmd = s.Resolved.PlayCmd
	s.Resolved.ThemesDir = s.File.ThemesDir
	s.Resolved.ThemeDir = filepath.Join(string(s.File.ThemesDir), string(s.File.ThemeRepo))
	s.AResolved.themeDir = s.Resolved.ThemeDir
	s.Resolved.ThemeRepo = s.File.ThemeRepo
	s.Resolved.CmdSounds = cloneMap(s.ThemeFile.CmdSounds)
	s.Resolved.NotFoundKey = notFoundKey

	overrides, ok := s.File.ThemeOverrides[s.File.ThemeRepo]
	if ok {
		for k, v := range overrides {
			s.Resolved.CmdSounds[k] = v
		}
	}

	if v, ok := s.Resolved.CmdSounds[notFoundKey]; ok {
		s.Resolved.NotFoundSound = v
		s.AResolved.notFoundSound = v
		delete(s.Resolved.CmdSounds, notFoundKey)
	}

	s.AResolved.cmdSounds = CmdSounds(s.Resolved.CmdSounds)
	s.AResolved.notFoundKey = notFoundKey

	return nil
}

func cloneMap(in map[string]string) map[string]string {
	out := make(map[string]string)

	for k, v := range in {
		out[k] = v
	}

	return out
}
