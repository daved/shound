package ccmd

import (
	"fmt"
	"io"

	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
)

type Identify struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func NewIdentify(out io.Writer, name string, cnf *config.Config) *Identify {
	fs := flagset.New(name)

	c := Identify{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	return &c
}

func (c *Identify) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Identify) HandleCommand() error { // NOTE: A
	args := c.fs.Args()
	argsLen := len(args)
	if argsLen == 0 {
		// TODO: A: return appropriate error
		return nil
	}

	sound := c.cnf.CmdSounds[args[len(args)-1]]
	fmt.Fprintln(c.out, sound)
	return nil
}
