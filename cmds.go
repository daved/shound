package main

import (
	"fmt"
	"os"

	"github.com/daved/flagset"
)

type CmdTop struct {
	cnf *Config
	fs  *flagset.FlagSet
}

func NewCmdTop(appName string, cnf *Config) *CmdTop {
	fs := flagset.New(appName)
	c := CmdTop{
		cnf: cnf,
		fs:  fs,
	}

	fs.Opt(&cnf.flags.help, "help|h", "help output", "")

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
}

func NewCmdExport(name string, cnf *Config) *CmdExport {
	fs := flagset.New(name)

	c := CmdExport{
		cnf: cnf,
		fs:  fs,
	}

	return &c
}

func (c *CmdExport) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *CmdExport) HandleCommand() error {
	// TODO: fill
	fmt.Println(c.cnf.file)
	fmt.Println()
	fmt.Fprint(os.Stdout, x)
	return nil
}
