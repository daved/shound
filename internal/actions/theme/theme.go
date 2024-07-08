package theme

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

type Theme struct {
	out io.Writer
	cnf *Config
}

func New(out io.Writer, cnf *Config) *Theme {
	return &Theme{
		out: out,
		cnf: cnf,
	}
}

func (a *Theme) Run(ctx context.Context) error {
	return fprintInfo(a.out, a.cnf.global)
}
