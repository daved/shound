package main

import (
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/internal/ccmds/export"
	"github.com/daved/shound/internal/ccmds/identify"
	"github.com/daved/shound/internal/ccmds/theme"
	"github.com/daved/shound/internal/ccmds/theme/info"
	"github.com/daved/shound/internal/ccmds/theme/install"
	"github.com/daved/shound/internal/ccmds/theme/list"
	"github.com/daved/shound/internal/ccmds/theme/set"
	"github.com/daved/shound/internal/ccmds/theme/uninstall"
	"github.com/daved/shound/internal/ccmds/top"
	"github.com/daved/shound/internal/config"
	"github.com/daved/shound/internal/thememgr"
)

func newCommand(appName string, out io.Writer, cnf *config.Config, tm *thememgr.ThemeMgr,
) (*clic.Clic, error) {
	cmd := top.New(out, appName, cnf).AsClic(
		export.New(out, "export", cnf).AsClic(),
		identify.New(out, "identify", cnf).AsClic(),
		theme.New(out, "theme", cnf).AsClic(
			install.New(out, "install", cnf).AsClic(),
			set.New(out, "set", cnf).AsClic(),
			list.New(out, "list", cnf, tm).AsClic(),
			info.New(out, "info", cnf).AsClic(),
			uninstall.New(out, "uninstall", cnf).AsClic(),
		),
	)

	return cmd, nil
}
