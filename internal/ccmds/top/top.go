package top

import (
	"errors"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type Top struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func NewTop(out io.Writer, name string, cnf *config.Config) *Top {
	fs := flagset.New(name)

	c := Top{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	fs.Opt(&cnf.User.Flags.Help, "help|h", "Print help output.")
	fs.Opt(&cnf.User.Flags.ConfFilePath, "conf", "Path to config file.")

	return &c
}

func (c *Top) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta()["SubRequired"] = true

	return cmd
}

func (c *Top) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Top) HandleCommand(cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	_ = ccmd.HandleHelpFlag(c.out, cmd, true)
	return errors.New("no subcommand")
}
