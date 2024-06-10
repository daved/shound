package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/daved/clic"

	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/thememgr"
)

func main() {
	var (
		appName        = "shound"
		debugEnvVarKey = "SHOUND_DEBUG"

		exitCode int
	)
	defer func() { os.Exit(exitCode) }()

	if _, debug := os.LookupEnv(debugEnvVarKey); debug {
		start := time.Now()
		var end time.Time
		defer func() {
			end = time.Now()
			fmt.Fprintln(os.Stderr, end.Sub(start))
		}()
	}

	if err := run(appName, os.Stdout, os.Args[1:]); err != nil {
		exitCode = 1
		if eerr, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = eerr.ExitCode()
		}

		fmt.Fprintf(os.Stderr, "%s: %v\n", appName, err)
		return
	}
}

func run(appName string, out io.Writer, args []string) error {
	var (
		configSubdir   = filepath.Join(".config", appName)
		configFileName = "config.yaml"
		themeFileName  = "shound.yaml"
	)

	defConfPath, err := defaultConfigurationFilePath(configSubdir, configFileName)
	if err != nil {
		return err
	}

	cnf, err := newConfig(defConfPath, themeFileName)
	if err != nil {
		return err
	}

	tm := thememgr.NewThemeMgr([]string{"flipflip.com/repo/test", "otherplace.com/project/this", "gityes.com/work/out"})

	cmd, err := newCommand(out, cnf, tm)
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
