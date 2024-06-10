package list

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
	"github.com/daved/shound/internal/thememgr"
)

type List struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
	tm  *thememgr.ThemeMgr
}

func New(out io.Writer, name string, cnf *config.Config, tm *thememgr.ThemeMgr) *List {
	fs := flagset.New(name)

	c := List{
		out: out,
		fs:  fs,
		cnf: cnf,
		tm:  tm,
	}

	return &c
}

func (c *List) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta()["CmdDesc"] = "List installed themes"

	return cmd
}

func (c *List) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *List) HandleCommand(ctx context.Context, cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	d := makeListData(c.cnf.ThemesDir, c.tm.Themes())

	return fprintList(c.out, d)
}
