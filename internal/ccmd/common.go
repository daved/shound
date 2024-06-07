package ccmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/daved/clic"
)

var ErrHelpFlag = errors.New("help requested")

func HandleHelpFlag(out io.Writer, cmd *clic.Clic, needsHelp bool) error {
	if needsHelp {
		fmt.Fprint(out, cmd.Usage())
		return ErrHelpFlag
	}

	return nil
}
