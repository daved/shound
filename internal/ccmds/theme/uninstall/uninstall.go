package uninstall

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

type Uninstall struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Uninstall {
	fs := flagset.New(name)

	c := Uninstall{
		out: out,
		fs:  fs,
		cnf: cnf,
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

	args := c.fs.Args()
	if len(args) == 0 {
		return errors.New("theme: uninstall: no theme repo")
	}
	arg := args[0]

	fmt.Fprintln(c.out, arg)

	return nil
}
