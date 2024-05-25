package ccmd

import (
	"fmt"
	"io"

	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
)

type Top struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func NewTop(out io.Writer, appName string, cnf *config.Config) *Top {
	fs := flagset.New(appName)

	c := Top{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	fs.Opt(&cnf.UserFlags.Help, "help|h", "print help output", "")
	fs.Opt(&cnf.UserFlags.ConfFilePath, "conf", "path to config file", "")

	return &c
}

func (c *Top) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Top) HandleCommand() error {
	if c.cnf.Help {
		fmt.Fprint(c.out, c.FlagSet().Help())
		return nil
	}
	return nil
}
