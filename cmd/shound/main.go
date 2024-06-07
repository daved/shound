package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/daved/clic"
	"github.com/daved/shound/internal/ccmd"
	"github.com/daved/shound/internal/config"
	"github.com/daved/shound/internal/tmpls"
)

var (
	appName        = "shound"
	configSubdir   = filepath.Join(".config", appName)
	configFileName = "config.yaml"
	themeFileName  = "shound.yaml"
	debugEnvVarKey = "SHOUND_DEBUG"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	if _, debug := os.LookupEnv(debugEnvVarKey); debug {
		start := time.Now()
		var end time.Time
		defer func() {
			end = time.Now()
			fmt.Fprintln(os.Stderr, end.Sub(start))
		}()
	}

	if err := run(os.Stdout, os.Args[1:]); err != nil {
		exitCode = 1
		if eerr, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = eerr.ExitCode()
		}

		fmt.Fprintf(os.Stderr, "%s: %v\n", appName, err)
		return
	}
}

func run(out io.Writer, args []string) error {
	defConfPath, err := defaultConfigurationFilePath()
	if err != nil {
		return err
	}

	cnf := config.NewConfig(defConfPath)

	cmd, err := newCommand(out, cnf)
	if err != nil {
		return err
	}

	if err := cmd.Parse(args); err != nil {
		if perr := (*clic.ParseError)(nil); errors.As(err, &perr) {
			fmt.Fprint(out, perr.Clic().Usage())
		}
		return err
	}

	cnfBytes, err := os.ReadFile(cnf.UserFlags.ConfFilePath)
	if err != nil {
		return err
	}

	if err := cnf.UserFile.InitFromYAML(cnfBytes); err != nil {
		return err
	}

	themeCnfBytes, err := os.ReadFile(cnf.UserFile.ThemePath(themeFileName))
	if err != nil {
		return err
	}

	if err := cnf.ThemeFile.InitFromYAML(themeCnfBytes); err != nil {
		return err
	}

	if err := cnf.Resolve(); err != nil {
		return err
	}

	if err := cmd.HandleCalled(); err != nil {
		if !errors.Is(err, ccmd.ErrHelpFlag) {
			return err
		}
	}

	return nil
}

func defaultConfigurationFilePath() (string, error) {
	eMsg := "default config file path: %v"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf(eMsg, err)
	}

	return filepath.Join(homeDir, configSubdir, configFileName), nil
}

func newCommand(out io.Writer, cnf *config.Config) (*clic.Clic, error) {
	eMsg := "new command: %v"

	ts, err := tmpls.NewTmpls()
	if err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	identify := ccmd.NewIdentify(out, "identify", cnf)
	cmdIdentify := clic.New(identify)

	export := ccmd.NewExport(out, ts, "export", cnf)
	cmdExport := clic.New(export)

	top := ccmd.NewTop(out, appName, cnf)
	return clic.New(top, cmdIdentify, cmdExport), nil
}
