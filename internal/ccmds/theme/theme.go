package ccmd

import (
	"errors"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
)

type Theme struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func NewTheme(out io.Writer, name string, cnf *config.Config) *Theme {
	fs := flagset.New(name)

	c := Theme{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	return &c
}

func (c *Theme) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta()["SubRequired"] = true

	return cmd
}

func (c *Theme) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Theme) HandleCommand(cmd *clic.Clic) error {
	if err := HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	_ = HandleHelpFlag(c.out, cmd, true)
	return errors.New("subcommand is required")
}
