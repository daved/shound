package install

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/internal/actions/theme/install"
	"github.com/daved/shound/internal/config"
)

type Install struct {
	fs  *flagset.FlagSet
	act *install.Install
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config, ta install.ThemeAdder) *Install {
	fs := flagset.New(name)

	act := install.New(out, install.NewConfig(cnf), ta)

	return &Install{
		fs:  fs,
		act: act,
		cnf: cnf,
	}
}

func (c *Install) AsClic(subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(cmd.NewHelpWrap(c.cnf, c), subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "Install a theme"
	cc.Meta[clic.MetaKeySubRequired] = true
	cc.Meta[clic.MetaKeyArgsHint] = "<theme_repo>"

	return cc
}

func (c *Install) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Install) HandleCommand(ctx context.Context) error {
	args := c.fs.Args()
	if len(args) == 0 {
		return fmt.Errorf("theme: install: %w", errors.New("no theme repo"))
	}
	themeRepo := args[0]

	return c.act.Run(ctx, themeRepo)
}
