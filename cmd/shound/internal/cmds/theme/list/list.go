package list

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/list"
)

type List struct {
	fs  *flagset.FlagSet
	act *list.List
	hr  cmd.HelpReporter
}

func New(out io.Writer, name string, cnf *config.Sourced, tp list.ThemesProvider) *List {
	fs := flagset.New(name)

	act := list.New(out, list.NewConfig(cnf.Resolved), tp)

	return &List{
		fs:  fs,
		act: act,
		hr:  cnf.AResolved,
	}
}

func (c *List) AsClic(subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.hr, c)

	cc := clic.New(h, subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "List installed themes"

	return cc
}

func (c *List) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *List) HandleCommand(ctx context.Context) error {
	return c.act.Run(ctx)
}
