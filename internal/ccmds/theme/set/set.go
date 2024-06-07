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

type Set struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Set {
	fs := flagset.New(name)

	c := Set{
		out: out,
		fs:  fs,
		cnf: cnf,
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

	args := c.fs.Args()
	if len(args) == 0 {
		return errors.New("theme: set: no theme repo")
	}
	arg := args[0]

	fmt.Fprintln(c.out, arg)

	return nil
}
