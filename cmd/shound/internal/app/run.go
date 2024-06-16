package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/daved/clic"
	"github.com/daved/shound/internal/ccmds/ccmd"
)

func Run(appName string, out io.Writer, args []string) error {
	var (
		configSubdir   = filepath.Join(".config", appName)
		configFileName = "config.yaml"
		themeFileName  = "shound.yaml"
	)

	// TODO: ensure config and cache dirs are present
	// TODO: support windows (config and cache dirs)

	defConfPath, err := defaultConfigurationFilePath(configSubdir, configFileName)
	if err != nil {
		return err
	}

	cnf, err := newConfig(defConfPath, themeFileName)
	if err != nil {
		return err
	}

	ti := newThemesInfo(out, cnf, themeFileName)

	cmd, err := newCommand(appName, out, cnf, ti)
	if err != nil {
		return err
	}

	if err := cmd.Parse(args); err != nil {
		if perr := (*clic.ParseError)(nil); errors.As(err, &perr) {
			fmt.Fprint(out, perr.Clic().Usage())
		}
		return err
	}

	if err := cmd.HandleCalled(context.Background()); err != nil {
		if !errors.Is(err, ccmd.ErrHelpFlag) {
			return err
		}
	}

	return nil
}
