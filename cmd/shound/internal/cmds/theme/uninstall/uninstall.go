package uninstall

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/uninstall"
)

type Uninstall struct {
	action *uninstall.Uninstall
	actCnf *uninstall.Config
}

func New(out io.Writer, td uninstall.ThemeDeleter, cnf *config.Sourced) *Uninstall {
	actCnf := uninstall.NewConfig(cnf.Resolved)

	return &Uninstall{
		action: uninstall.New(out, td, actCnf),
		actCnf: actCnf,
	}
}

func (c *Uninstall) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(c, name, subs...)
	cc.Description = "Uninstall a theme"

	cc.Operand(&c.actCnf.ThemeRepo, true, "theme_repo", "")

	return cc
}

func (c *Uninstall) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
