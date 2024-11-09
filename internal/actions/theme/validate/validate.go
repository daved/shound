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
	Theme string
	Hash  string
}

func NewConfig() *Config {
	return &Config{}
}

type Validate struct {
	out io.Writer
	tv  ThemeValidator
	cnf *Config
}

func New(out io.Writer, tv ThemeValidator, cnf *Config) *Validate {
	return &Validate{
		out: out,
		tv:  tv,
		cnf: cnf,
	}
}

func (a *Validate) Run(ctx context.Context) error {
	eMsg := "theme: validate: %w"

	if a.cnf.IsDir {
		if a.cnf.Theme == "" { // ignore hash
			return fmt.Errorf(eMsg, errors.New("no theme directory"))
		}

		if err := a.tv.ValidateThemeDir(a.cnf.Theme); err != nil {
			return fmt.Errorf(eMsg, err)
		}

		return nil
	}

	if a.cnf.Theme == "" {
		return fmt.Errorf(eMsg, errors.New("no theme repo"))
	}

	if err := a.tv.ValidateThemeRemote(a.cnf.Theme, a.cnf.Hash); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
