package uninstall

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/daved/shound/internal/actions/acterr"
	"github.com/daved/shound/internal/config"
)

type ThemeDeleter interface {
	IsThemeInstalled(string) (bool, error)
	DeleteTheme(string) error
}

type Config struct {
	global *config.Config
}

func NewConfig(global *config.Config) *Config {
	return &Config{
		global: global,
	}
}

type Uninstall struct {
	out io.Writer
	cnf *Config
	td  ThemeDeleter
}

func New(out io.Writer, cnf *Config, td ThemeDeleter) *Uninstall {
	return &Uninstall{
		out: out,
		cnf: cnf,
		td:  td,
	}
}

func (a *Uninstall) Run(ctx context.Context, themeRepo string) error {
	eMsg := "theme: uninstall: %w"

	if themeRepo == a.cnf.global.ThemeRepo {
		return fmt.Errorf(eMsg, errors.New("theme repo same as current theme"))
	}

	isInstalled, err := a.td.IsThemeInstalled(themeRepo)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if !isInstalled {
		return acterr.NewNotInstalledError(themeRepo)
	}

	if err := a.td.DeleteTheme(themeRepo); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
