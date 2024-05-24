package ccmd

import (
	"os"

	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
	"github.com/daved/shound/internal/tmpls"
)

type Export struct {
	cnf *config.Config
	fs  *flagset.FlagSet
	ts  *tmpls.Tmpls
}

func NewExport(name string, cnf *config.Config, ts *tmpls.Tmpls) *Export {
	fs := flagset.New(name)

	c := Export{
		cnf: cnf,
		fs:  fs,
		ts:  ts,
	}

	return &c
}

func (c *Export) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Export) HandleCommand() error {
	d := tmpls.AliasesData{
		PlayCmd:    c.cnf.PlayCmd,
		SoundDir:   string(c.cnf.SoundDir),
		CmdsSounds: c.cnf.CmdSounds,
		NoCmdSound: c.cnf.NoCmdSound,
	}

	return c.ts.Aliases(os.Stdout, d)
}
