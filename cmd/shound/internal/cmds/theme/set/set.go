package set

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/set"
)

type Set struct {
	action *set.Set
	actCnf *set.Config
	appCnf *config.Sourced
}

func New(out io.Writer, ts set.ThemeSetter, cnf *config.Sourced) *Set {
	actCnf := set.NewConfig()

	return &Set{
		action: set.New(out, ts, actCnf),
		actCnf: actCnf,
		appCnf: cnf,
	}
}

func (c *Set) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.appCnf.AResolved, c)

	cc := clic.New(h, name, subs...)
	cc.SubRequired = true

	cc.ArgSet.Arg(&c.actCnf.ThemeRepo, false, "theme_repo", "")

	cc.UsageConfig.CmdDesc = "Set the current theme"
	cc.UsageConfig.ArgsHint = "<theme_repo[@{<hash>|latest}]>"

	return cc
}

func (c *Set) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
