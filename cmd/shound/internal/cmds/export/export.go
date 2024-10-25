package export

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/export"
)

type Export struct {
	fs  *flagset.FlagSet
	act *export.Export
	hr  cmd.HelpReporter
}

func New(out io.Writer, name string, cnf *config.Sourced) *Export {
	fs := flagset.New(name)

	act := export.New(out, cnf.AResolved)

	return &Export{
		fs:  fs,
		act: act,
		hr:  cnf.AResolved,
	}
}

func (c *Export) AsClic(subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.hr, c)

	cc := clic.New(h, subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "Print code for a shell to evaluate"

	return cc
}

func (c *Export) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Export) HandleCommand(ctx context.Context) error {
	return c.act.Run(ctx)
}
