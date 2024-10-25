package app

import (
	"io"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/cmds/export"
	"github.com/daved/shound/cmd/shound/internal/cmds/identify"
	"github.com/daved/shound/cmd/shound/internal/cmds/root"
	"github.com/daved/shound/cmd/shound/internal/cmds/theme"
	"github.com/daved/shound/cmd/shound/internal/cmds/theme/install"
	"github.com/daved/shound/cmd/shound/internal/cmds/theme/list"
	"github.com/daved/shound/cmd/shound/internal/cmds/theme/set"
	"github.com/daved/shound/cmd/shound/internal/cmds/theme/uninstall"
	"github.com/daved/shound/cmd/shound/internal/cmds/theme/validate"
	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/cmd/shound/internal/themesmgr"
)

type (
	conf      = config.Sourced
	themesMgr = themesmgr.ThemesMgr
)

func newCommand(appName string, out io.Writer, cnf *conf, tm *themesMgr) (*clic.Clic, error) {
	cmd := root.New(appName, cnf).AsClic(
		export.New(out, "export", cnf).AsClic(),
		identify.New(out, "identify", cnf).AsClic(),
		theme.New(out, "theme", cnf).AsClic(
			install.New(out, "install", cnf, tm).AsClic(),
			set.New(out, "set", cnf, tm).AsClic(),
			list.New(out, "list", cnf, tm).AsClic(),
			uninstall.New(out, "uninstall", cnf, tm).AsClic(),
			validate.New(out, "validate", cnf, tm).AsClic(),
		),
	)

	return cmd, nil
}
