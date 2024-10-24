package config

import (
	"bufio"
	"bytes"
	"fmt"

	"gopkg.in/yaml.v3"
)

type ThemeOverrides map[string]map[string]string // map[ThemeRepo]map[CommandName]SoundFile

type File struct {
	Bypass         bool           `yaml:"Bypass"`
	PlayCmd        string         `yaml:"PlayCmd"`
	ThemesDir      string         `yaml:"ThemesDir"`
	ThemeRepo      string         `yaml:"ThemeRepo"`
	ThemeOverrides ThemeOverrides `yaml:"CmdSoundsOverrides"`
}

func (f *File) InitFromYAML(data []byte) error {
	eMsg := "config: file: init from yaml: %w"

	if err := yaml.Unmarshal(data, f); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (f *File) validate() error {

	eMsg := "file: validate: %s"

	if f.PlayCmd == "" {
		return fmt.Errorf(eMsg, `play command missing (PlayCmd: "pw-play")`)
	}

	if f.ThemeRepo == "" {
		return fmt.Errorf(eMsg, `theme repo missing (ThemeRepo: "star_trek")`)
	}

	return nil
}

const (
	FileFieldThemeRepo = "ThemeRepo"
)

func SetFileDataField(data []byte, name string, value any) ([]byte, error) {
	eMsg := "config: set file data field: %w"

	var out []byte
	var found bool

	buf := bytes.NewBuffer(data)
	sc := bufio.NewScanner(buf)
	for sc.Scan() {
		bs := sc.Bytes()
		n := []byte(name)
		if bytes.HasPrefix(bs, n) {
			found = true

			_, comment, _ := bytes.Cut(bs, []byte("#"))
			bs = append(n, []byte(fmt.Sprintf(`: "%v"`, value))...)
			if len(comment) > 0 {
				bs = append(bs, ' ', '#')
				bs = append(bs, comment...)
			}

		}
		out = append(out, bs...)
		out = append(out, '\n')
	}

	if !found {
		return nil, fmt.Errorf(eMsg, fmt.Errorf("%q not found", name))
	}

	return out, nil
}
