package ccmd

import (
	"fmt"
	"os"

	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
	"github.com/daved/shound/internal/tmpls"
)

type CmdTop struct {
	cnf *config.Config
	fs  *flagset.FlagSet

	ConfFilePath string
}

func NewCmdTop(appName string, cnf *config.Config, defaultConfFile string) *CmdTop {
	fs := flagset.New(appName)
	c := CmdTop{
		cnf:          cnf,
		fs:           fs,
		ConfFilePath: defaultConfFile,
	}

	fs.Opt(&cnf.Flags.Help, "help|h", "print help output", "")
	fs.Opt(&c.ConfFilePath, "conf", "path to config file", "")

	return &c
}

func (c *CmdTop) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *CmdTop) HandleCommand() error {
	if c.cnf.Help {
		fmt.Fprint(os.Stdout, c.FlagSet().Help())
		return nil
	}
	return nil
}

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
