package install

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/daved/shound/internal/actions/acterr"
)

type ThemeAdder interface {
	IsThemeInstalled(string) (bool, error)
	AddTheme(theme, hash string) error
}

type Config struct {
	ThemeRepo string
	ThemeHash string
}

func NewConfig() *Config {
	return &Config{}
}

type Install struct {
	out io.Writer
	cnf *Config
	ta  ThemeAdder
}

func New(out io.Writer, ta ThemeAdder, cnf *Config) *Install {
	return &Install{
		out: out,
		cnf: cnf,
		ta:  ta,
	}
}

func (a *Install) Run(ctx context.Context) error {
	eMsg := "theme: install: %w"

	if a.cnf.ThemeRepo == "" {
		return fmt.Errorf(eMsg, errors.New("no theme repo"))
	}

	isInstalled, err := a.ta.IsThemeInstalled(a.cnf.ThemeRepo)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if isInstalled {
		return acterr.NewAlreadyInstalledError(a.cnf.ThemeRepo)
	}

	// TODO: if version different, checkout correct version

	// TODO: handle versions on initial clone
	if err := a.ta.AddTheme(a.cnf.ThemeRepo, a.cnf.ThemeHash); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
