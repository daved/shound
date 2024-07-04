package theme

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type Theme struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Theme {
	fs := flagset.New(name)

	c := Theme{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	return &c
}

func (c *Theme) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta[clic.MetaKeyCmdDesc] = "Show info about the current theme"

	return cmd
}

func (c *Theme) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Theme) HandleCommand(ctx context.Context, cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	return fprintInfo(c.out, c.cnf)
}
