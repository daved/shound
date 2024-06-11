package uninstall

import (
	"context"
	"errors"
	"fmt"
	"io"
	"slices"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type ThemeDeleter interface {
	Themes() ([]string, error)
	DeleteTheme(string) error
}

type Uninstall struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
	td  ThemeDeleter
}

func New(out io.Writer, name string, cnf *config.Config, td ThemeDeleter) *Uninstall {
	fs := flagset.New(name)

	c := Uninstall{
		out: out,
		fs:  fs,
		cnf: cnf,
		td:  td,
	}

	return &c
}

func (c *Uninstall) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta()["CmdDesc"] = "Uninstall a theme"
	cmd.Meta()["SubRequired"] = true
	cmd.Meta()["ArgsHint"] = "<theme_repo>"

	return cmd
}

func (c *Uninstall) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Uninstall) HandleCommand(ctx context.Context, cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	eMsg := "theme: uninstall: %w"

	args := c.fs.Args()
	if len(args) == 0 {
		return fmt.Errorf(eMsg, errors.New("no theme repo"))
	}
	arg := args[0]

	if arg == c.cnf.ThemeRepo {
		return fmt.Errorf(eMsg, errors.New("theme repo same as current theme"))
	}

	themes, err := c.td.Themes()
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if !slices.Contains(themes, arg) {
		err := fmt.Errorf("%q is not a valid installed theme", arg)
		return fmt.Errorf(eMsg, err)
	}

	if err := c.td.DeleteTheme(arg); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
