package ccmd

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
)

type Identify struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config

	playCmd bool
}

func NewIdentify(out io.Writer, name string, cnf *config.Config) *Identify {
	fs := flagset.New(name)

	c := Identify{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	fs.Opt(&c.playCmd, "playcmd", "prefix identified sound with play command string", "")

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
	arg := args[len(args)-1]

	sound, ok := c.cnf.CmdSounds[arg]
	if !ok && arg == c.cnf.NotFoundKey {
		sound = c.cnf.NotFoundSound
	}

	soundDir := filepath.Join(string(c.cnf.SoundCache), string(c.cnf.Theme))
	soundPath := filepath.Join(soundDir, sound)

	if c.playCmd {
		fmt.Fprintf(c.out, "%s %s\n", c.cnf.PlayCmd, soundPath)
		return nil
	}

	fmt.Fprintln(c.out, soundPath)
	return nil
}
