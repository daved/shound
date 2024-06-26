package app

import (
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/internal/ccmds/export"
	"github.com/daved/shound/internal/ccmds/identify"
	"github.com/daved/shound/internal/ccmds/theme"
	"github.com/daved/shound/internal/ccmds/theme/install"
	"github.com/daved/shound/internal/ccmds/theme/list"
	"github.com/daved/shound/internal/ccmds/theme/set"
	"github.com/daved/shound/internal/ccmds/theme/uninstall"
	"github.com/daved/shound/internal/ccmds/top"
	"github.com/daved/shound/internal/config"
)

func newCommand(appName string, out io.Writer, cnf *config.Config, ti *themesInfo,
) (*clic.Clic, error) {
	cmd := top.New(out, appName, cnf).AsClic(
		export.New(out, "export", cnf).AsClic(),
		identify.New(out, "identify", cnf).AsClic(),
		theme.New(out, "theme", cnf).AsClic(
			install.New(out, "install", cnf, ti).AsClic(),
			set.New(out, "set", cnf, ti).AsClic(),
			list.New(out, "list", cnf, ti).AsClic(),
			uninstall.New(out, "uninstall", cnf, ti).AsClic(),
			// TODO: add validate subcmd
		),
	)

	return cmd, nil
}
