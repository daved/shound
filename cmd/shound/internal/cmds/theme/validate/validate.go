package validate

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/theme/validate"
)

type Validate struct {
	fs  *flagset.FlagSet
	act *validate.Validate
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config, tv validate.ThemeValidator) *Validate {
	fs := flagset.New(name)

	actCnf := validate.NewConfig()
	fs.Opt(&actCnf.IsDir, "dir", "Use theme_repo arg as a directory (hash will be ignored).")

	act := validate.New(out, actCnf, tv)

	return &Validate{
		fs:  fs,
		act: act,
		cnf: cnf,
	}
}

func (c *Validate) AsClic(subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.cnf.Resolved, c)

	cc := clic.New(h, subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "Validate a theme"
	cc.Meta[clic.MetaKeySubRequired] = true
	cc.Meta[clic.MetaKeyArgsHint] = "<theme_repo> [hash]"

	return cc
}

func (c *Validate) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Validate) HandleCommand(ctx context.Context) error {
	themeRepo := c.fs.Arg(0)
	themeHash := c.fs.Arg(1)

	return c.act.Run(ctx, themeRepo, themeHash)
}
