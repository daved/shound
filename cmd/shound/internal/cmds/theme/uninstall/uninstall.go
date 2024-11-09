package uninstall

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/uninstall"
)

type Uninstall struct {
	action *uninstall.Uninstall
	actCnf *uninstall.Config
	appCnf *config.Sourced
}

func New(out io.Writer, td uninstall.ThemeDeleter, cnf *config.Sourced) *Uninstall {
	actCnf := uninstall.NewConfig(cnf.Resolved)

	return &Uninstall{
		action: uninstall.New(out, td, actCnf),
		actCnf: actCnf,
		appCnf: cnf,
	}
}

func (c *Uninstall) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.appCnf.AResolved, c)

	cc := clic.New(h, name, subs...)
	cc.SubRequired = true

	cc.ArgSet.Arg(&c.actCnf.ThemeRepo, true, "theme_repo", "")

	cc.UsageConfig.CmdDesc = "Uninstall a theme"
	cc.UsageConfig.ArgsHint = "<theme_repo>"

	return cc
}

func (c *Uninstall) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
