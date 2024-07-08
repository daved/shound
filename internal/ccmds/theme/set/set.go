package set

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type ThemeSetter interface {
	IsThemeInstalled(string) (bool, error)
	SetTheme(string) error
}

type Set struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
	ts  ThemeSetter
}

func New(out io.Writer, name string, cnf *config.Config, ts ThemeSetter) *Set {
	fs := flagset.New(name)

	c := Set{
		out: out,
		fs:  fs,
		cnf: cnf,
		ts:  ts,
	}

	return &c
}

func (c *Set) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta[clic.MetaKeyCmdDesc] = "Set the current theme"
	cmd.Meta[clic.MetaKeySubRequired] = true
	cmd.Meta[clic.MetaKeyArgsHint] = "<theme_repo[@{<hash>|latest}]>"

	return cmd
}

func (c *Set) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Set) HandleCommand(ctx context.Context) error {
	if c.cnf.Help {
		return ccmd.NewUsageError(ccmd.ErrHelpFlag)
	}

	eMsg := "theme: set: %w"

	args := c.fs.Args()
	if len(args) == 0 {
		return fmt.Errorf(eMsg, errors.New("no theme repo"))
	}
	theme := args[0]

	isInstalled, err := c.ts.IsThemeInstalled(theme)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if !isInstalled {
		return ccmd.NewNotInstalledError(theme)
	}

	if err := c.ts.SetTheme(theme); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
