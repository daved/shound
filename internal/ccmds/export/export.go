package export

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type Export struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Export {
	fs := flagset.New(name)

	c := Export{
		out: out,
		cnf: cnf,
		fs:  fs,
	}

	return &c
}

func (c *Export) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta()["CmdDesc"] = "Print code for a shell to evaluate"

	return cmd
}

func (c *Export) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Export) HandleCommand(ctx context.Context, cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	aliases := c.cnf.CmdSounds.CmdList()
	d := makeAliasesData(c.cnf.NotFoundKey, c.cnf.NotFoundSound, aliases)

	return fprintAliases(c.out, d)
}
