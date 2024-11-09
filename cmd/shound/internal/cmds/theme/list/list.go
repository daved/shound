package list

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/list"
)

type List struct {
	action *list.List
	appCnf *config.Sourced
}

func New(out io.Writer, tp list.ThemesProvider, cnf *config.Sourced) *List {
	return &List{
		action: list.New(out, list.NewConfig(cnf.Resolved), tp),
		appCnf: cnf,
	}
}

func (c *List) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.appCnf.AResolved, c)

	cc := clic.New(h, name, subs...)
	cc.UsageConfig.CmdDesc = "List installed themes"

	return cc
}

func (c *List) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
