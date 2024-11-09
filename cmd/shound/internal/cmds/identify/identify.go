package identify

import (
	"context"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/internal/actions/identify"
)

type Identify struct {
	action *identify.Identify
	actCnf *identify.Config
	appCnf *config.Sourced
}

func New(out io.Writer, cnf *config.Sourced) *Identify {
	actCnf := identify.NewConfig()

	return &Identify{
		action: identify.New(out, cnf.AResolved, actCnf),
		actCnf: actCnf,
		appCnf: cnf,
	}
}

func (c *Identify) AsClic(name string, subs ...*clic.Clic) *clic.Clic {
	h := cmd.NewHelpWrap(c.appCnf.AResolved, c)

	cc := clic.New(h, name, subs...)
	cc.FlagSet.Opt(&c.actCnf.PlayCmd, "playcmd", "Prefix identified sound with play command string.")
	cc.ArgSet.Arg(&c.actCnf.CmdName, false, "command_name", "")
	cc.UsageConfig.CmdDesc = "Print file associated with the provided command"
	// cc.Meta[clic.MetaKeyArgsHint] = "<command_name>"

	return cc
}

func (c *Identify) HandleCommand(ctx context.Context) error {
	return c.action.Run(ctx)
}
