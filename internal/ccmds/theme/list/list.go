package ccmd

import (
	"fmt"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
)

type ThemeList struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func NewThemeList(out io.Writer, name string, cnf *config.Config) *ThemeList {
	fs := flagset.New(name)

	c := ThemeList{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	return &c
}

func (c *ThemeList) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta()["CmdDesc"] = "List installed themes"

	return cmd
}

func (c *ThemeList) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *ThemeList) HandleCommand(cmd *clic.Clic) error {
	if err := HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	fmt.Fprintln(c.out, "listing...")

	return nil
}
