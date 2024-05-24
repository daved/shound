package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/daved/clic"
	"github.com/daved/shound/internal/ccmd"
	"github.com/daved/shound/internal/config"
	"github.com/daved/shound/internal/tmpls"
)

var (
	appName        = "shound"
	configSubdir   = filepath.Join(".config", appName)
	configFileName = "config.toml"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	if err := run(os.Args[1:]); err != nil {
		if eerr, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = eerr.ExitCode()
		}
		fmt.Printf("%s: %v\n", appName, err)
	}
}

func run(args []string) error { // TODO: handle errors
	defConfPath, err := defaultConfigurationFilePath()
	if err != nil {
		return err
	}

	cnf := config.NewConfig(defConfPath)

	ts, err := tmpls.NewTmpls()
	if err != nil {
		return err
	}

	top := ccmd.NewTop(appName, cnf)
	export := ccmd.NewExport("export", cnf, ts)

	cmdExport := clic.New(export)
	cmd := clic.New(top, cmdExport)

	if err := cmd.Parse(args); err != nil {
		if perr := (*clic.ParseError)(nil); errors.As(err, &perr) {
			fmt.Println(perr.Handler().FlagSet().Help())
		}
		return err
	}

	cnfHandle, err := os.Open(cnf.UserFlags.ConfFilePath)
	if err != nil {
		return err
	}
	if err := cnf.UserFile.InitFromTOML(cnfHandle); err != nil {
		return err
	}

	if err := cnf.Resolve(); err != nil {
		return err
	}

	if err := cmd.HandleCommand(); err != nil {
		return err
	}

	return nil
}

func defaultConfigurationFilePath() (string, error) { // TODO: handle errors
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configSubdir, configFileName), nil
}
