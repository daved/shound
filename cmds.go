package main

import (
	"os"

	"github.com/daved/flagset"
)

type CmdTop struct {
	cnf *Config
	fs  *flagset.FlagSet

	confFilePath string
}

func NewCmdTop(appName string, cnf *Config, defaultConfFile string) *CmdTop {
	fs := flagset.New(appName)
	c := CmdTop{
		cnf:          cnf,
		fs:           fs,
		confFilePath: defaultConfFile,
	}

	fs.Opt(&cnf.flags.help, "help|h", "print help output", "")
	fs.Opt(&c.confFilePath, "conf", "path to config file", "")

	return &c
}

func (c *CmdTop) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *CmdTop) HandleCommand() error {
	// TODO: should print help?
	return nil
}

type CmdExport struct {
	cnf *Config
	fs  *flagset.FlagSet
	ts  *Tmpls
}

func NewCmdExport(name string, cnf *Config, ts *Tmpls) *CmdExport {
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
	d := AliasesData{
		PlayCmd:    c.cnf.playCmd,
		SoundDir:   string(c.cnf.soundDir),
		CmdsSounds: c.cnf.cmdsSounds,
	}

	return c.ts.Aliases(os.Stdout, d)
}
