package cmd

import (
	"context"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/config"
)

type HelpWrap struct {
	h   clic.Handler
	cnf *config.Config
}

func NewHelpWrap(cnf *config.Config, h clic.Handler) *HelpWrap {
	return &HelpWrap{
		h:   h,
		cnf: cnf,
	}
}

func (c *HelpWrap) FlagSet() *flagset.FlagSet {
	return c.h.FlagSet()
}

func (c *HelpWrap) HandleCommand(ctx context.Context) error {
	if c.cnf.Help {
		return NewUsageError(ErrHelpFlag)
	}

	return c.h.HandleCommand(ctx)
}
