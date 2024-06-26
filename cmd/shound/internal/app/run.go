package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/daved/clic"
	"github.com/daved/shound/internal/ccmds/ccmd"
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
