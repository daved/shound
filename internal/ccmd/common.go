package ccmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/internal/config"
)

var ErrHelpFlag = errors.New("help requested")

func HandleHelpFlag(out io.Writer, cnf *config.Config, cmd *clic.Clic) error {
	if cnf.Help {
		fmt.Fprint(out, cmd.Usage())
		return ErrHelpFlag
	}

	return nil
}
