package ccmd

import (
	"os"

	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
	"github.com/daved/shound/internal/tmpls"
)

type CmdExport struct {
	cnf *config.Config
	fs  *flagset.FlagSet
	ts  *tmpls.Tmpls
}

func NewCmdExport(name string, cnf *config.Config, ts *tmpls.Tmpls) *CmdExport {
	fs := flagset.New(name)

	c := CmdExport{
		cnf: cnf,
		fs:  fs,
		ts:  ts,
	}

	return &c
}

func (c *CmdExport) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *CmdExport) HandleCommand() error {
	d := tmpls.AliasesData{
		PlayCmd:    c.cnf.PlayCmd,
		SoundDir:   string(c.cnf.SoundDir),
		CmdsSounds: c.cnf.CmdsSounds,
		NoCmdSound: c.cnf.NoCmdSound,
	}

	return c.ts.Aliases(os.Stdout, d)
}
