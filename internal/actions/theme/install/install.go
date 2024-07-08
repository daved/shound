package install

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/daved/shound/internal/actions/acterr"
	"github.com/daved/shound/internal/config"
)

type ThemeAdder interface {
	IsThemeInstalled(string) (bool, error)
	AddTheme(string) error
}

type Config struct {
	global *config.Config
}

func NewConfig(global *config.Config) *Config {
	return &Config{
		global: global,
	}
}

type Install struct {
	out io.Writer
	cnf *Config
	ta  ThemeAdder
}

func New(out io.Writer, cnf *Config, ta ThemeAdder) *Install {
	return &Install{
		out: out,
		cnf: cnf,
		ta:  ta,
	}
}

func (a *Install) Run(ctx context.Context, themeRepo string) error {
	eMsg := "theme: install: %w"

	if themeRepo == "" {
		return fmt.Errorf(eMsg, errors.New("no theme repo"))
	}

	isInstalled, err := a.ta.IsThemeInstalled(themeRepo)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if isInstalled {
		return acterr.NewAlreadyInstalledError(themeRepo)
	}

	// TODO: if version different, checkout correct version

	// TODO: handle versions on initial clone
	if err := a.ta.AddTheme(themeRepo); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
