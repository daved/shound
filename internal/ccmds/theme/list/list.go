package list

import (
	"context"
	"fmt"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type ThemesProvider interface {
	Themes() ([]string, error)
}

type List struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
	tp  ThemesProvider
}

func New(out io.Writer, name string, cnf *config.Config, tp ThemesProvider) *List {
	fs := flagset.New(name)

	c := List{
		out: out,
		fs:  fs,
		cnf: cnf,
		tp:  tp,
	}

	return &c
}

func (c *List) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta[clic.MetaKeyCmdDesc] = "List installed themes"

	return cmd
}

func (c *List) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *List) HandleCommand(ctx context.Context, cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	themes, err := c.tp.Themes()
	if err != nil {
		return fmt.Errorf("theme: list: %w", err)
	}

	d := makeListData(c.cnf.ThemesDir, themes)

	return fprintList(c.out, d)
}
