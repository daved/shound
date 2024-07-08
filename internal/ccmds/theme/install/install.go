package install

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

type ThemeAdder interface {
	IsThemeInstalled(string) (bool, error)
	AddTheme(string) error
}

type Install struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
	ta  ThemeAdder
}

func New(out io.Writer, name string, cnf *config.Config, ta ThemeAdder) *Install {
	fs := flagset.New(name)

	c := Install{
		out: out,
		fs:  fs,
		cnf: cnf,
		ta:  ta,
	}

	return &c
}

func (c *Install) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta[clic.MetaKeyCmdDesc] = "Install a theme"
	cmd.Meta[clic.MetaKeySubRequired] = true
	cmd.Meta[clic.MetaKeyArgsHint] = "<theme_repo>"

	return cmd
}

func (c *Install) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Install) HandleCommand(ctx context.Context) error {
	if c.cnf.Help {
		return ccmd.NewUsageError(ccmd.ErrHelpFlag)
	}

	eMsg := "theme: install: %w"

	args := c.fs.Args()
	if len(args) == 0 {
		return fmt.Errorf(eMsg, errors.New("no theme repo"))
	}
	theme := args[0]

	isInstalled, err := c.ta.IsThemeInstalled(theme)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if isInstalled {
		return ccmd.NewAlreadyInstalledError(theme)
	}

	// TODO: if version different, checkout correct version

	// TODO: handle versions on initial clone
	if err := c.ta.AddTheme(theme); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
