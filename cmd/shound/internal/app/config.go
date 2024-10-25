package app

import (
	"fmt"
	"path/filepath"

	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/cmd/shound/internal/fs"
)

func newConfig(fs fs.FS, defConfPath, defThemesDirPath, themeFileName string) (*config.Sourced, error) {
	eMsg := "new config: %w"

	// should take NewFile(defConfPath, defThemesDirPath).InitFromYAML([]byte)
	// should also take NewThemeFile().InitFromYAML([]byte)
	// sourced should be reworked:
	// consider if validation stays on sourced, or on individual file types
	// sourced should return a "resolved" config
	cnf := config.NewSourced(defConfPath, defThemesDirPath)

	cnfBytes, err := fs.ReadFile(cnf.Flags.ConfFilePath)
	if err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.File.InitFromYAML(cnfBytes); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	themePath := filepath.Join(cnf.File.ThemesDir, cnf.File.ThemeRepo, themeFileName)

	themeCnfBytes, err := fs.ReadFile(themePath)
	if err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.ThemeFile.InitFromYAML(themeCnfBytes); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	if err := cnf.ValidateFiles(); err != nil {
		return nil, fmt.Errorf(eMsg, err)
	}

	return cnf, nil
}
