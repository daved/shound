package theme

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme"
)

type Theme struct {
	action *theme.Theme
	actCnf *theme.Config
	hr     cmd.HelpReporter
}

func New(out io.Writer, cnf *config.Sourced) *Theme {
	actCnf := theme.NewConfig(cnf.Resolved)

	c := Theme{
		action: theme.New(out, actCnf),
		hr:     cnf.AResolved,
	}

	return &c
}

func (c *Theme) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.hr, c)

	cc := clic.New(h, name, subs...)
	cc.UsageConfig.CmdDesc = "Show info about the current theme"

	return cc
}

func (c *Theme) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
