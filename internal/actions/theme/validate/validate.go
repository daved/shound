package validate

import (
	"context"
	"errors"
	"fmt"
	"io"
)

type ThemeValidator interface {
	ValidateThemeRemote(repo, hash string) error
	ValidateThemeDir(string) error
}

type Config struct {
	IsDir bool
}

func NewConfig() *Config {
	return &Config{}
}

type Validate struct {
	out io.Writer
	cnf *Config
	tv  ThemeValidator
}

func New(out io.Writer, cnf *Config, tv ThemeValidator) *Validate {
	return &Validate{
		out: out,
		cnf: cnf,
		tv:  tv,
	}
}

func (a *Validate) Run(ctx context.Context, theme, hash string) error {
	eMsg := "theme: validate: %w"

	if a.cnf.IsDir {
		if theme == "" { // ignore hash
			return fmt.Errorf(eMsg, errors.New("no theme directory"))
		}

		if err := a.tv.ValidateThemeDir(theme); err != nil {
			return fmt.Errorf(eMsg, err)
		}

		return nil
	}

	if theme == "" {
		return fmt.Errorf(eMsg, errors.New("no theme repo"))
	}

	if err := a.tv.ValidateThemeRemote(theme, hash); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
