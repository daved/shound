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
	appCnf *config.Sourced
}

func New(cnf *config.Sourced) *Root {
	return &Root{
		appCnf: cnf,
	}
}

func (c *Root) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(c, name, subs...)
	cc.SubRequired = true

	cc.Flag(&c.appCnf.Flags.ConfFilePath, "conf", "Path to config file.")

	helpOpt := cc.FlagRecursive(&c.appCnf.Flags.Help, "help|h", "Print help output.")
	helpOpt.Meta[flagset.MetaKeyTypeHint] = ""
	helpOpt.Meta[flagset.MetaKeyDefaultHint] = ""

	return cc
}

func (c *Root) HandleCommand(ctx context.Context) error {
	return cmd.NewUsageError(errors.New("no subcommand"))
}
