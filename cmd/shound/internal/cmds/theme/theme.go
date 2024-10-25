package theme

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme"
)

type Theme struct {
	fs  *flagset.FlagSet
	act *theme.Theme
	hr  cmd.HelpReporter
}

func New(out io.Writer, name string, cnf *config.Sourced) *Theme {
	fs := flagset.New(name)

	act := theme.New(out, theme.NewConfig(cnf.Resolved))

	c := Theme{
		fs:  fs,
		act: act,
		hr:  cnf.AResolved,
	}

	return &c
}

func (c *Theme) AsClic(subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.hr, c)

	cc := clic.New(h, subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "Show info about the current theme"

	return cc
}

func (c *Theme) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Theme) HandleCommand(ctx context.Context) error {
	return c.act.Run(ctx)
}
