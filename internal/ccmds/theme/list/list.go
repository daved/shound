package list

import (
	"fmt"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type List struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *List {
	fs := flagset.New(name)

	c := List{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	return &c
}

func (c *List) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta()["CmdDesc"] = "List installed themes"

	return cmd
}

func (c *List) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *List) HandleCommand(cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	fmt.Fprintln(c.out, "listing...")

	return nil
}
