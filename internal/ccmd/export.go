package ccmd

import (
	"io"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
	"github.com/daved/shound/internal/tmpls"
)

type Export struct {
	out io.Writer
	ts  *tmpls.Tmpls

	fs  *flagset.FlagSet
	cnf *config.Config
}

func NewExport(out io.Writer, ts *tmpls.Tmpls, name string, cnf *config.Config) *Export {
	fs := flagset.New(name)

	c := Export{
		out: out,
		ts:  ts,
		cnf: cnf,
		fs:  fs,
	}

	return &c
}

func (c *Export) AsClic(subs ...*clic.Clic) *clic.Clic {
	return clic.New(c, subs...)
}

func (c *Export) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Export) HandleCommand(cmd *clic.Clic) error {
	if err := HandleHelpFlag(c.out, c.cnf, cmd); err != nil {
		return err
	}

	aliases := make([]string, 0, len(c.cnf.CmdSounds))
	for alias := range c.cnf.CmdSounds {
		aliases = append(aliases, alias)
	}

	d := tmpls.MakeAliasesData(c.cnf.NotFoundKey, c.cnf.NotFoundSound, aliases)

	return c.ts.Aliases(c.out, d)
}
