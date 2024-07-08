package identify

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/daved/shound/internal/config"
)

type Config struct {
	global *config.Config

	PlayCmd bool
}

func NewConfig(global *config.Config) *Config {
	return &Config{
		global: global,
	}
}

type Identify struct {
	out io.Writer
	cnf *Config
}

func New(out io.Writer, cnf *Config) *Identify {
	return &Identify{
		out: out,
		cnf: cnf,
	}
}

func (a *Identify) Run(ctx context.Context, cmdName string) error {
	if cmdName == "" {
		return errors.New("identify: no command name")
	}

	gCnf := a.cnf.global

	sound, ok := gCnf.CmdSounds[cmdName]
	if !ok && cmdName == gCnf.NotFoundKey {
		sound = gCnf.NotFoundSound
	}

	soundPath := filepath.Join(gCnf.ThemeDir, sound)

	if a.cnf.PlayCmd {
		fmt.Fprintf(a.out, "%s %s\n", gCnf.PlayCmd, soundPath)
		return nil
	}

	fmt.Fprintln(a.out, soundPath)
	return nil
}
