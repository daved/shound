package list

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/list"
)

type List struct {
	action *list.List
}

func New(out io.Writer, tp list.ThemesProvider, cnf *config.Sourced) *List {
	return &List{
		action: list.New(out, list.NewConfig(cnf.Resolved), tp),
	}
}

func (c *List) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(c, name, subs...)
	cc.Description = "List installed themes"

	return cc
}

func (c *List) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
