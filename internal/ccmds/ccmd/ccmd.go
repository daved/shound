package ccmd

import (
	"fmt"
	"io"

	"github.com/daved/clic"
)

func HandleHelpFlag(out io.Writer, cmd *clic.Clic, needsHelp bool) error {
	if needsHelp {
		fmt.Fprint(out, cmd.Usage())
		return ErrHelpFlag
	}

	return nil
}
