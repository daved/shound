package identify

import (
	"context"
	"errors"
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/internal/actions/identify"
	"github.com/daved/shound/internal/config"
)

type Identify struct {
	fs  *flagset.FlagSet
	act *identify.Identify
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Identify {
	fs := flagset.New(name)

	actCnf := identify.NewConfig(cnf)
	fs.Opt(&actCnf.PlayCmd, "playcmd", "Prefix identified sound with play command string.")

	act := identify.New(out, actCnf)

	return &Identify{
		fs:  fs,
		act: act,
		cnf: cnf,
	}
}

func (c *Identify) AsClic(subs ...*clic.Clic) *clic.Clic {
	cc := clic.New(cmd.NewHelpWrap(c.cnf, c), subs...)
	cc.Meta[clic.MetaKeyCmdDesc] = "Print file associated with the provided command"
	cc.Meta[clic.MetaKeyArgsHint] = "<command_name>"

	return cc
}

func (c *Identify) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Identify) HandleCommand(ctx context.Context) error {
	args := c.fs.Args()
	if len(args) == 0 {
		return errors.New("identify: no command name")
	}
	cmdName := args[0]

	return c.act.Run(ctx, cmdName)
}
