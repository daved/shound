package set

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/daved/shound/internal/actions/acterr"
)

type ThemeSetter interface {
	IsThemeInstalled(string) (bool, error)
	SetTheme(string) error
}

type Config struct {
	ThemeRepo string
}

func NewConfig() *Config {
	return &Config{}
}

type Set struct {
	out io.Writer
	ts  ThemeSetter
	cnf *Config
}

func New(out io.Writer, ts ThemeSetter, cnf *Config) *Set {
	return &Set{
		out: out,
		ts:  ts,
		cnf: cnf,
	}
}

func (a *Set) Run(ctx context.Context) error {
	eMsg := "theme: set: %w"
	themeRepo := a.cnf.ThemeRepo

	if themeRepo == "" {
		return fmt.Errorf(eMsg, errors.New("no theme repo"))
	}

	isInstalled, err := a.ts.IsThemeInstalled(themeRepo)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if !isInstalled {
		return acterr.NewNotInstalledError(themeRepo)
	}

	if err := a.ts.SetTheme(themeRepo); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
