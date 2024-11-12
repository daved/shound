package export

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/export"
)

type Export struct {
	action *export.Export
	appCnf *config.Sourced
}

func New(out io.Writer, cnf *config.Sourced) *Export {
	return &Export{
		action: export.New(out, cnf.AResolved),
		appCnf: cnf,
	}
}

func (c *Export) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(c, name, subs...)
	cc.UsageConfig.CmdDesc = "Print code for a shell to evaluate"

	return cc
}

func (c *Export) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
