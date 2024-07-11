package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/daved/clic"
	"github.com/daved/shound/cmd/shound/internal/cmds/cmd"
)

func Run(appName string, out io.Writer, args []string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	var (
		configFileName   = "config.yaml"
		defConfigDirPath = filepath.Join(homeDir, ".config", appName)
		defConfigPath    = filepath.Join(defConfigDirPath, configFileName)
		themeFileName    = "shound.yaml"
		defThemesDirPath = filepath.Join(homeDir, ".cache", appName, "themes")
	)

	// TODO: support windows (config and cache dirs)

	if err = fsEnsureDirsExist(defConfigDirPath, defThemesDirPath); err != nil {
		return err
	}

	cnf, err := newConfig(defConfigPath, defThemesDirPath, themeFileName)
	if err != nil {
		return err
	}

	ti := NewThemesInfo(appName, out, cnf, themeFileName)

	cc, err := newCommand(appName, out, cnf, ti)
	if err != nil {
		return err
	}

	if err := cc.Parse(args); err != nil {
		if perr := (*clic.ParseError)(nil); errors.As(err, &perr) {
			fmt.Fprint(out, perr.Clic().Usage())
		}
		return err
	}

	if err := cnf.Resolve(); err != nil {
		return err
	}

	if err := cc.Handle(context.Background()); err != nil {
		if uerr := (*cmd.UsageError)(nil); errors.As(err, &uerr) {
			fmt.Fprint(out, cc.Called().Usage())
		}

		if !errors.Is(err, cmd.ErrHelpFlag) {
			return err
		}
	}

	return nil
}
