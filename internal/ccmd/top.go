package ccmd

import (
	"fmt"
	"os"

	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
)

type Top struct {
	cnf *config.Config
	fs  *flagset.FlagSet

	ConfFilePath string
}

func NewTop(appName string, cnf *config.Config, defaultConfFile string) *Top {
	fs := flagset.New(appName)
	c := Top{
		cnf:          cnf,
		fs:           fs,
		ConfFilePath: defaultConfFile,
	}

	fs.Opt(&cnf.Flags.Help, "help|h", "print help output", "")
	fs.Opt(&c.ConfFilePath, "conf", "path to config file", "")

	return &c
}

func (c *Top) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Top) HandleCommand() error {
	if c.cnf.Help {
		fmt.Fprint(os.Stdout, c.FlagSet().Help())
		return nil
	}
	return nil
}
