package identify

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
)

type Identify struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config

	playCmd bool
}

func New(out io.Writer, name string, cnf *config.Config) *Identify {
	fs := flagset.New(name)

	c := Identify{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	fs.Opt(&c.playCmd, "playcmd", "Prefix identified sound with play command string.")

	return &c
}

func (c *Identify) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta[clic.MetaKeyCmdDesc] = "Print file associated with the provided command"
	cmd.Meta[clic.MetaKeyArgsHint] = "<command_name>"

	return cmd
}

func (c *Identify) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Identify) HandleCommand(ctx context.Context, cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	args := c.fs.Args()
	if len(args) == 0 {
		return errors.New("identify: no command name")
	}
	arg := args[0]

	sound, ok := c.cnf.CmdSounds[arg]
	if !ok && arg == c.cnf.NotFoundKey {
		sound = c.cnf.NotFoundSound
	}

	soundPath := filepath.Join(c.cnf.ThemeDir, sound)

	if c.playCmd {
		fmt.Fprintf(c.out, "%s %s\n", c.cnf.PlayCmd, soundPath)
		return nil
	}

	fmt.Fprintln(c.out, soundPath)
	return nil
}
