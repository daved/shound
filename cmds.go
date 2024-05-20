package main

import (
	"fmt"

	"github.com/daved/flagset"
)

type Cnf struct {
	help bool
	info string
}

type CmdTop struct {
	*Cnf

	fs *flagset.FlagSet
}

func NewCmdTop(appName string, cnf *Cnf) *CmdTop {
	fs := flagset.New(appName)
	c := CmdTop{
		Cnf: cnf,
		fs:  fs,
	}

	fs.Opt(&c.help, "help|h", "help output", "")
	fs.Opt(&c.info, "info|i", "info test", "")

	return &c
}

func (c *CmdTop) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *CmdTop) HandleCommand() error {
	fmt.Println("info", c.info)
	return nil
}

type CmdExport struct {
	cnf   *Cnf
	extra string

	fs *flagset.FlagSet
}

func NewCmdExport(name string, cnf *Cnf) *CmdExport {
	fs := flagset.New(name)

	c := CmdExport{
		cnf: cnf,
		fs:  fs,
	}

	fs.Opt(&c.extra, "extra|e", "extra info", "")

	return &c
}

func (c *CmdExport) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *CmdExport) HandleCommand() error {
	fmt.Println("export info", c.cnf.info)
	fmt.Println("export extra", c.extra)
	return nil
}
