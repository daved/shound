package cmd

import (
	"context"

	"github.com/daved/clic"
	"github.com/daved/flagset"
)

type HelpReporter interface {
	Help() bool
}

type HelpWrap struct {
	next clic.Handler
	hr   HelpReporter
}

func NewHelpWrap(hr HelpReporter, next clic.Handler) *HelpWrap {
	return &HelpWrap{
		next: next,
		hr:   hr,
	}
}

func (c *HelpWrap) FlagSet() *flagset.FlagSet {
	return c.next.FlagSet()
}

func (c *HelpWrap) HandleCommand(ctx context.Context) error {
	if c.hr.Help() {
		return NewUsageError(ErrHelpFlag)
	}

	return c.next.HandleCommand(ctx)
}
