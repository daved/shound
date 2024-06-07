package install

import (
	"errors"
	"fmt"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type ThemeInstall struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func NewThemeInstall(out io.Writer, name string, cnf *config.Config) *ThemeInstall {
	fs := flagset.New(name)

	c := ThemeInstall{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	return &c
}

func (c *ThemeInstall) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta()["CmdDesc"] = "Install a theme"
	cmd.Meta()["SubRequired"] = true
	cmd.Meta()["ArgsHint"] = "<theme_repo>"

	return cmd
}

func (c *ThemeInstall) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *ThemeInstall) HandleCommand(cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	args := c.fs.Args()
	if len(args) == 0 {
		return errors.New("theme: install: no theme repo")
	}
	arg := args[0]

	fmt.Fprintln(c.out, arg)

	return nil
}
