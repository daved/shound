package set

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/internal/actions/theme/set"
	"github.com/daved/shound/internal/config"
)

type Set struct {
	fs  *flagset.FlagSet
	act *set.Set
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config, ts set.ThemeSetter) *Set {
	fs := flagset.New(name)

	act := set.New(out, ts)

	return &Set{
		fs:  fs,
		act: act,
		cnf: cnf,
	}
}

func (c *Set) AsClic(subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(cmd.NewHelpWrap(c.cnf, c), subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "Set the current theme"
	cc.Meta[clic.MetaKeySubRequired] = true
	cc.Meta[clic.MetaKeyArgsHint] = "<theme_repo[@{<hash>|latest}]>"

	return cc
}

func (c *Set) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Set) HandleCommand(ctx context.Context) error {
	args := c.fs.Args()
	if len(args) == 0 {
		return fmt.Errorf("theme: set: %w", errors.New("no theme repo"))
	}
	themeRepo := args[0]

	return c.act.Run(ctx, themeRepo)
}
