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
	cc.Description = "Validate a theme"

	cc.Operand(&c.actCnf.Theme, true, "theme_repo", "Theme repo to validate.")
	cc.Operand(&c.actCnf.Hash, false, "hash", "Theme repo hash (falls back to latest).")

	cc.Flag(&c.actCnf.IsDir, "dir", "Use theme_repo arg as a directory (hash will be ignored).")

	return cc
}

func (c *Validate) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
