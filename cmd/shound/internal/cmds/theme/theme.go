package theme

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme"
)

type Theme struct {
	action *theme.Theme
}

func New(out io.Writer, cnf *config.Sourced) *Theme {
	actCnf := theme.NewConfig(cnf.Resolved)

	c := Theme{
		action: theme.New(out, actCnf),
	}

	return &c
}

func (c *Theme) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(c, name, subs...)
	cc.Description = "Show info about the current theme"

	return cc
}

func (c *Theme) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
