package export

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/internal/actions/export"
	"github.com/daved/shound/internal/config"
)

type Export struct {
	fs  *flagset.FlagSet
	act *export.Export
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Export {
	fs := flagset.New(name)

	act := export.New(out, export.NewConfig(cnf))

	return &Export{
		fs:  fs,
		act: act,
		cnf: cnf,
	}
}

func (c *Export) AsClic(subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.cnf, c)

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
