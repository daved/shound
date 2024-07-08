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

type Set struct {
	out io.Writer
	ts  ThemeSetter
}

func New(out io.Writer, ts ThemeSetter) *Set {
	return &Set{
		out: out,
		ts:  ts,
	}
}

func (a *Set) Run(ctx context.Context, themeRepo string) error {
	eMsg := "theme: set: %w"

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
