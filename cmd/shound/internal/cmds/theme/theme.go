package theme

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/internal/actions/theme"
	"github.com/daved/shound/internal/config"
)

type Theme struct {
	fs  *flagset.FlagSet
	act *theme.Theme
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Theme {
	fs := flagset.New(name)

	act := theme.New(out, theme.NewConfig(cnf))

	c := Theme{
		fs:  fs,
		act: act,
		cnf: cnf,
	}

	return &c
}

func (c *Theme) AsClic(subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(cmd.NewHelpWrap(c.cnf, c), subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "Show info about the current theme"

	return cc
}

func (c *Theme) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Theme) HandleCommand(ctx context.Context) error {
	return c.act.Run(ctx)
}
