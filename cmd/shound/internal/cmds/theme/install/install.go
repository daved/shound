package install

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/install"
)

type Install struct {
	fs  *flagset.FlagSet
	act *install.Install
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config, ta install.ThemeAdder) *Install {
	fs := flagset.New(name)

	act := install.New(out, install.NewConfig(cnf.Resolved), ta)

	return &Install{
		fs:  fs,
		act: act,
		cnf: cnf,
	}
}

func (c *Install) AsClic(subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.cnf.Resolved, c)

	cc := clic.New(h, subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "Install a theme"
	cc.Meta[clic.MetaKeySubRequired] = true
	cc.Meta[clic.MetaKeyArgsHint] = "<theme_repo>"

	return cc
}

func (c *Install) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Install) HandleCommand(ctx context.Context) error {
	themeRepo := c.fs.Arg(0)
	themeHash := c.fs.Arg(1)

	return c.act.Run(ctx, themeRepo, themeHash)
}
