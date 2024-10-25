package export

import (
	"context"
	"io"
)

type CmdSoundsReporter interface {
	CmdList() []string
	NotFoundKey() string
	NotFoundSound() string
}

type Export struct {
	out io.Writer
	csr CmdSoundsReporter
}

func New(out io.Writer, csr CmdSoundsReporter) *Export {
	return &Export{
		out: out,
		csr: csr,
	}
}

func (a *Export) Run(ctx context.Context) error {
	aliases := a.csr.CmdList()
	d := makeAliasesData(a.csr.NotFoundKey(), a.csr.NotFoundSound(), aliases)

	return fprintAliases(a.out, d)
}
