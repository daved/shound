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
	"github.com/daved/shound/cmd/shound/internal/fs/opfs"
	"github.com/daved/shound/cmd/shound/internal/themesmgr"
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
	fs := opfs.New()

	if err = ensureDirsExist(fs, defConfigDirPath, defThemesDirPath); err != nil {
		return err
	}

	cnf, err := newConfig(fs, defConfigPath, defThemesDirPath, themeFileName)
	if err != nil {
		return err
	}

	tm := themesmgr.New(fs, out, appName, themeFileName, cnf)

	cc, err := newCommand(appName, out, cnf, tm)
	if err != nil {
		return err
	}

	err = cc.Parse(args)
	if err != nil {
		if cerr := (*clic.Error)(nil); errors.As(err, &cerr) {
			fmt.Fprintln(out, cerr.Clic().ResolvedCmd().Usage())
		}
		return err
	}
	if cnf.Flags.Help {
		fmt.Fprintln(out, cc.ResolvedCmd().Usage())
		return nil
	}

	if err := cnf.Resolve(); err != nil {
		return err
	}

	if err := cc.HandleResolvedCmd(context.Background()); err != nil {
		if uerr := (*cmd.UsageError)(nil); errors.As(err, &uerr) {
			fmt.Fprint(out, cc.ResolvedCmd().Usage())
		}
	}

	return nil
}
