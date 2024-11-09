package config

type Config struct {
	cmdSounds     CmdSounds
	notFoundKey   string
	notFoundSound string
	help          bool
	themeDir      string
	playCmd       string
}

func (c *Config) CmdSounds() map[string]string {
	return c.cmdSounds
}

func (c *Config) CmdList() []string {
	return c.cmdSounds.CmdList()
}

func (c *Config) NotFoundKey() string {
	return c.notFoundKey
}

func (c *Config) NotFoundSound() string {
	return c.notFoundSound
}

func (c *Config) Help() bool {
	return c.help
}

func (c *Config) ThemeDir() string {
	return c.themeDir
}

func (c *Config) PlayCmd() string {
	return c.playCmd
}
