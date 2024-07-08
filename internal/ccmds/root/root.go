package root

import (
	"context"
	"errors"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type Root struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Root {
	fs := flagset.New(name)

	c := Root{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	fs.Opt(&cnf.User.Flags.Help, "help|h", "Print help output.")
	fs.Opt(&cnf.User.Flags.ConfFilePath, "conf", "Path to config file.")

	return &c
}

func (c *Root) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta[clic.MetaKeySubRequired] = true

	return cmd
}

func (c *Root) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Root) HandleCommand(ctx context.Context) error {
	if c.cnf.Help {
		return ccmd.NewUsageError(ccmd.ErrHelpFlag)
	}

	return ccmd.NewUsageError(errors.New("no subcommand"))
}
