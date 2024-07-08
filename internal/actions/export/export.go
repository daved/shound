package export

import (
	"context"
	"io"

	"github.com/daved/shound/internal/config"
)

type Config struct {
	global *config.Config
}

func NewConfig(global *config.Config) *Config {
	return &Config{
		global: global,
	}
}

type Export struct {
	out io.Writer
	cnf *Config
}

func New(out io.Writer, cnf *Config) *Export {
	return &Export{
		out: out,
		cnf: cnf,
	}
}

func (a *Export) Run(ctx context.Context) error {
	gCnf := a.cnf.global

	aliases := gCnf.CmdSounds.CmdList()
	d := makeAliasesData(gCnf.NotFoundKey, gCnf.NotFoundSound, aliases)

	return fprintAliases(a.out, d)
}
