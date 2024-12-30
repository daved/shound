package install

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/internal/actions/theme/install"
)

type Install struct {
	action *install.Install
	actCnf *install.Config
}

func New(out io.Writer, ta install.ThemeAdder) *Install {
	actCnf := install.NewConfig()

	return &Install{
		action: install.New(out, ta, actCnf),
		actCnf: actCnf,
	}
}

func (c *Install) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(c, name, subs...)
	cc.Description = "Install a theme"

	cc.Operand(&c.actCnf.ThemeRepo, true, "theme_repo", "")
	cc.Operand(&c.actCnf.ThemeHash, false, "hash", "")

	return cc
}

func (c *Install) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
