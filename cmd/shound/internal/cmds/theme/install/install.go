package install

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/install"
)

type Install struct {
	action *install.Install
	actCnf *install.Config
	appCnf *config.Sourced
}

func New(out io.Writer, ta install.ThemeAdder, cnf *config.Sourced) *Install {
	actCnf := install.NewConfig()

	return &Install{
		action: install.New(out, ta, actCnf),
		actCnf: actCnf,
		appCnf: cnf,
	}
}

func (c *Install) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.appCnf.AResolved, c)

	cc := clic.New(h, name, subs...)
	cc.SubRequired = true

	cc.ArgSet.Arg(&c.actCnf.ThemeRepo, true, "theme_repo", "")
	cc.ArgSet.Arg(&c.actCnf.ThemeHash, false, "theme_hash", "")

	cc.UsageConfig.CmdDesc = "Install a theme"
	cc.UsageConfig.ArgsHint = "<theme_repo> [hash]"

	return cc
}

func (c *Install) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
