package config

type Config struct {
	Help         bool
	ConfFilePath string

	Bypass    bool
	PlayCmd   string
	ThemesDir string
	ThemeDir  string
	ThemeRepo string

	CmdSounds     CmdSounds
	NotFoundKey   string
	NotFoundSound string
}

type CmdSounds map[string]string // map[CommandName]SoundFile

func (css CmdSounds) CmdList() []string {
	cmds := make([]string, 0, len(css))
	for cmd := range css {
		cmds = append(cmds, cmd)
	}
	return cmds
}
