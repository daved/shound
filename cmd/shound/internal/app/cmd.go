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
	cmd := root.New(cnf).AsClic(appName,
		export.New(out, cnf).AsClic("export"),
		identify.New(out, cnf).AsClic("identify"),
		theme.New(out, cnf).AsClic("theme",
			install.New(out, tm, cnf).AsClic("install"),
			set.New(out, tm, cnf).AsClic("set"),
			list.New(out, tm, cnf).AsClic("list"),
			uninstall.New(out, tm, cnf).AsClic("uninstall"),
			validate.New(out, tm, cnf).AsClic("validate"),
		),
	)

	return cmd, nil
}
