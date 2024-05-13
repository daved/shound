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

type CmdHook struct {
	cnf   *Cnf
	extra string

	fs *flagset.FlagSet
}

func NewCmdHook(name string, cnf *Cnf) *CmdHook {
	fs := flagset.New(name)

	c := CmdHook{
		cnf: cnf,
		fs:  fs,
	}

	fs.Opt(&c.extra, "extra|e", "extra info", "")

	return &c
}

func (c *CmdHook) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *CmdHook) HandleCommand() error {
	fmt.Println("hook info", c.cnf.info)
	fmt.Println("hook extra", c.extra)
	return nil
}

type CmdOther struct {
	cnf  *Cnf
	more string

	fs *flagset.FlagSet
}

func NewCmdOther(name string, cnf *Cnf) *CmdOther {
	fs := flagset.New(name)

	c := CmdOther{
		cnf: cnf,
		fs:  fs,
	}

	fs.Opt(&c.more, "more|m", "more info", "")

	return &c
}

func (c *CmdOther) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *CmdOther) HandleCommand() error {
	fmt.Println("other info", c.cnf.info)
	fmt.Println("other more", c.more)
	return nil
}
