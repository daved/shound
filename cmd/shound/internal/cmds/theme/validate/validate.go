package validate

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/internal/actions/theme/validate"
)

type Validate struct {
	action *validate.Validate
	actCnf *validate.Config
}

func New(out io.Writer, tv validate.ThemeValidator) *Validate {
	actCnf := validate.NewConfig()

	return &Validate{
		action: validate.New(out, tv, actCnf),
		actCnf: actCnf,
	}
}

func (c *Validate) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(c, name, subs...)

	cc.Flag(&c.actCnf.IsDir, "dir", "Use theme_repo arg as a directory (hash will be ignored).")

	cc.Arg(&c.actCnf.Theme, true, "theme_repo", "Theme repo to validate.")
	cc.Arg(&c.actCnf.Hash, false, "hash", "Specific theme repo hash (falls back to latest).")

	cc.UsageConfig.CmdDesc = "Validate a theme"

	return cc
}

func (c *Validate) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
