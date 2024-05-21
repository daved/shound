package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/daved/clic"
)

var (
	appName      = "shound"
	configSubdir = filepath.Join(".config", appName)
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

func run(args []string) error { // TODO: error handling
	cnf := NewConfig()

	ts, err := NewTmpls()
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	defaultConfFile := filepath.Join(homeDir, configSubdir, "config.toml")

	top := NewCmdTop(appName, cnf, defaultConfFile)
	export := NewCmdExport("export", cnf, ts)

	cmdExport := clic.New(export)
	cmd := clic.New(top, cmdExport)

	if err := cmd.Parse(args); err != nil {
		if perr := (*clic.ParseError)(nil); errors.As(err, &perr) {
			fmt.Println(perr.FlagSet().Help())
		}
		return err
	}

	fmt.Println(top.confFilePath)
	cnfHandle, err := os.Open(top.confFilePath)
	if err != nil {
		return err
	}
	if err := cnf.file.initFromTOML(cnfHandle); err != nil {
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
