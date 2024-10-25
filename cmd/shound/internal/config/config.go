package config

type Config struct {
	cmdSounds     CmdSounds
	notFoundKey   string
	notFoundSound string
	help          bool
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
