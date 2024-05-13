package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/daved/clic"
)

var (
	appName = "shound"
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

func run(args []string) error {
	cnf := &Cnf{}

	top := NewCmdTop(appName, cnf)
	hook := NewCmdHook("hook", cnf)
	other := NewCmdOther("other", cnf)

	cmdHook := clic.New(hook)
	cmdOther := clic.New(other)

	cmd := clic.New(top, cmdHook, cmdOther)
	if err := cmd.Parse(args); err != nil {
		if perr := (*clic.ParseError)(nil); errors.As(err, &perr) {
			fmt.Println(perr.FlagSet().Help())
		}
		return err
	}

	return cmd.HandleCommand()
}
