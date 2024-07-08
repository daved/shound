package uninstall

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/internal/actions/theme/uninstall"
	"github.com/daved/shound/internal/config"
)

type Uninstall struct {
	fs  *flagset.FlagSet
	act *uninstall.Uninstall
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config, td uninstall.ThemeDeleter) *Uninstall {
	fs := flagset.New(name)

	act := uninstall.New(out, uninstall.NewConfig(cnf), td)

	return &Uninstall{
		fs:  fs,
		act: act,
		cnf: cnf,
	}
}

func (c *Uninstall) AsClic(subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(cmd.NewHelpWrap(c.cnf, c), subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "Uninstall a theme"
	cc.Meta[clic.MetaKeySubRequired] = true
	cc.Meta[clic.MetaKeyArgsHint] = "<theme_repo>"

	return cc
}

func (c *Uninstall) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Uninstall) HandleCommand(ctx context.Context) error {
	themeRepo := c.fs.Arg(0)

	return c.act.Run(ctx, themeRepo)
}
