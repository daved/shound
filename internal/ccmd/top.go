package ccmd

import (
	"errors"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
)

type Top struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func NewTop(out io.Writer, appName string, cnf *config.Config) *Top {
	fs := flagset.New(appName)

	c := Top{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	fs.Opt(&cnf.UserFlags.Help, "help|h", "Print help output.")
	fs.Opt(&cnf.UserFlags.ConfFilePath, "conf", "Path to config file.")

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
	if err := HandleHelpFlag(c.out, c.cnf, cmd); err != nil {
		return err
	}

	c.cnf.Help = true
	_ = HandleHelpFlag(c.out, c.cnf, cmd)
	return errors.New("subcommand is required")
}
