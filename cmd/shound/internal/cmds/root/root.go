package root

import (
	"context"
	"errors"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
)

type Root struct {
	fs *flagset.FlagSet
	hr cmd.HelpReporter
}

func New(name string, cnf *config.Sourced) *Root {
	fs := flagset.New(name)

	c := Root{
		fs: fs,
		hr: cnf.AResolved,
	}

	fs.Opt(&cnf.Flags.Help, "help|h", "Print help output.")
	fs.Opt(&cnf.Flags.ConfFilePath, "conf", "Path to config file.")

	return &c
}

func (c *Root) AsClic(subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.hr, c)

	cc := clic.New(h, subs...)
	cc.Meta[clic.MetaKeySubRequired] = true

	return cc
}

func (c *Root) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Root) HandleCommand(ctx context.Context) error {
	return cmd.NewUsageError(errors.New("no subcommand"))
}
