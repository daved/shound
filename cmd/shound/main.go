package main

import (
	"fmt"
	"os"
	"time"

	"github.com/daved/shound/cmd/shound/internal/app"
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

	if err := app.Run(appName, os.Stdout, os.Args[1:]); err != nil {
		exitCode = 1
		if eerr, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = eerr.ExitCode()
		}

		fmt.Fprintf(os.Stderr, "%s: %v\n", appName, err)
		return
	}
}
