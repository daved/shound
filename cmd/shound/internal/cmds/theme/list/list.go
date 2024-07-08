package list

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/internal/actions/theme/list"
	"github.com/daved/shound/internal/config"
)

type List struct {
	fs  *flagset.FlagSet
	act *list.List
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config, tp list.ThemesProvider) *List {
	fs := flagset.New(name)

	act := list.New(out, list.NewConfig(cnf), tp)

	return &List{
		fs:  fs,
		act: act,
		cnf: cnf,
	}
}

func (c *List) AsClic(subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(cmd.NewHelpWrap(c.cnf, c), subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "List installed themes"

	return cc
}

func (c *List) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *List) HandleCommand(ctx context.Context) error {
	return c.act.Run(ctx)
}
