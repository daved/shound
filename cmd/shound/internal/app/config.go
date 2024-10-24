package app

import (
	"fmt"
	"path/filepath"

	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/cmd/shound/internal/fs"
)

func newConfig(fs fs.FS, defConfPath, defThemesDirPath, themeFileName string) (*config.Config, error) {
	eMsg := "new config: %w"

	cnf := config.New(defConfPath, defThemesDirPath)

	cnfBytes, err := fs.ReadFile(cnf.User.Flags.ConfFilePath)
	if err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.User.File.InitFromYAML(cnfBytes); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	themePath := filepath.Join(cnf.User.File.ThemesDir, cnf.User.File.ThemeRepo, themeFileName)

	themeCnfBytes, err := fs.ReadFile(themePath)
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
