package identify

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
)

type CmdSoundsReporter interface {
	CmdSounds() map[string]string
	NotFoundKey() string
	NotFoundSound() string
	ThemeDir() string
	PlayCmd() string
}

type Config struct {
	PlayCmd bool
	CmdName string
}

func NewConfig() *Config {
	return &Config{}
}

type Identify struct {
	out io.Writer
	csr CmdSoundsReporter
	cnf *Config
}

func New(out io.Writer, csr CmdSoundsReporter, cnf *Config) *Identify {
	return &Identify{
		out: out,
		csr: csr,
		cnf: cnf,
	}
}

func (a *Identify) Run(ctx context.Context) error {
	if a.cnf.CmdName == "" {
		return errors.New("identify: no command name")
	}

	csr := a.csr
	cmdName := a.cnf.CmdName
	playCmd := a.cnf.PlayCmd

	sound, ok := csr.CmdSounds()[cmdName]
	if !ok && cmdName == csr.NotFoundKey() {
		sound = csr.NotFoundSound()
	}

	soundPath := filepath.Join(csr.ThemeDir(), sound)

	if playCmd {
		fmt.Fprintf(a.out, "%s %s\n", csr.PlayCmd(), soundPath)
		return nil
	}

	fmt.Fprintln(a.out, soundPath)
	return nil
}
