package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/daved/shound/internal/config"
)

func newConfig(defConfPath, defThemesDirPath, themeFileName string) (*config.Config, error) {
	eMsg := "new config: %w"

	cnf := config.New(defConfPath, defThemesDirPath)

	cnfBytes, err := os.ReadFile(cnf.User.Flags.ConfFilePath)
	if err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.User.File.InitFromYAML(cnfBytes); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	themePath := filepath.Join(cnf.User.File.ThemesDir, cnf.User.File.ThemeRepo, themeFileName)

	themeCnfBytes, err := os.ReadFile(themePath)
	if err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.User.ThemeFile.InitFromYAML(themeCnfBytes); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.ValidateFiles(); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	return cnf, nil
}
