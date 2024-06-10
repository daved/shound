package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/daved/shound/internal/config"
)

func defaultConfigurationFilePath(cnfSubdir, cnfFileName string) (string, error) {
	eMsg := "default config file path: %v"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf(eMsg, err)
	}

	return filepath.Join(homeDir, cnfSubdir, cnfFileName), nil
}

func newConfig(defConfPath, themeFileName string) (*config.Config, error) {
	eMsg := "new config: %w"

	cnf := config.New(defConfPath)

	cnfBytes, err := os.ReadFile(cnf.User.Flags.ConfFilePath)
	if err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.User.File.InitFromYAML(cnfBytes); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	themeCnfBytes, err := os.ReadFile(cnf.User.File.ThemePath(themeFileName))
	if err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.User.ThemeFile.InitFromYAML(themeCnfBytes); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.Resolve(); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	return cnf, nil
}
