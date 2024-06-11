package set

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

type ThemeSetter interface {
	Themes() ([]string, error)
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
	cmd.Meta()["CmdDesc"] = "Set the current theme"
	cmd.Meta()["SubRequired"] = true
	cmd.Meta()["ArgsHint"] = "<theme_repo[@{<hash>|latest}]>"

	return cmd
}

func (c *Set) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Set) HandleCommand(ctx context.Context, cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	eMsg := "theme: set: %w"

	args := c.fs.Args()
	if len(args) == 0 {
		return fmt.Errorf(eMsg, errors.New("no theme repo"))
	}
	arg := args[0]

	themes, err := c.ts.Themes()
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if !slices.Contains(themes, arg) {
		err := fmt.Errorf("%q is not a valid installed theme", arg)
		return fmt.Errorf(eMsg, err)
	}

	if err := c.ts.SetTheme(arg); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}
