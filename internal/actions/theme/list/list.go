package list

import (
	"context"
	"fmt"
	"io"

	"github.com/daved/shound/internal/config"
)

type ThemesProvider interface {
	Themes() ([]string, error)
}

type Config struct {
	global *config.Config
}

func NewConfig(global *config.Config) *Config {
	return &Config{
		global: global,
	}
}

type List struct {
	out io.Writer
	cnf *Config
	tp  ThemesProvider
}

func New(out io.Writer, cnf *Config, tp ThemesProvider) *List {
	return &List{
		out: out,
		cnf: cnf,
		tp:  tp,
	}
}

func (a *List) Run(ctx context.Context) error {
	themes, err := a.tp.Themes()
	if err != nil {
		return fmt.Errorf("theme: list: %w", err)
	}

	d := makeListData(a.cnf.global.ThemesDir, themes)

	return fprintList(a.out, d)
}
