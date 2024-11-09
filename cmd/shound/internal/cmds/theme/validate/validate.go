package validate

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/validate"
)

type Validate struct {
	action *validate.Validate
	actCnf *validate.Config
	appCnf *config.Sourced
}

func New(out io.Writer, tv validate.ThemeValidator, cnf *config.Sourced) *Validate {
	actCnf := validate.NewConfig()

	return &Validate{
		action: validate.New(out, tv, actCnf),
		actCnf: actCnf,
		appCnf: cnf,
	}
}

func (c *Validate) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.appCnf.AResolved, c)

	cc := clic.New(h, name, subs...)
	cc.SubRequired = true

	cc.FlagSet.Opt(&c.actCnf.IsDir, "dir", "Use theme_repo arg as a directory (hash will be ignored).")

	cc.ArgSet.Arg(&c.actCnf.Theme, true, "theme_repo", "Theme repo to validate.")
	cc.ArgSet.Arg(&c.actCnf.Hash, false, "hash", "Specific theme repo hash (falls back to latest).")

	cc.UsageConfig.CmdDesc = "Validate a theme"
	cc.UsageConfig.ArgsHint = "<theme_repo> [hash]"

	return cc
}

func (c *Validate) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
