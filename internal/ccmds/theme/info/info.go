package info

import (
	"fmt"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type Info struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Info {
	fs := flagset.New(name)

	c := Info{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	return &c
}

func (c *Info) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta()["CmdDesc"] = "Show info about the current theme"

	return cmd
}

func (c *Info) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Info) HandleCommand(cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	fmt.Fprintln(c.out, "info...")

	return nil
}
