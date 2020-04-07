package transform

const (
	_defIndexLeaves = true
	_defValidate    = true
)

type Config struct {
	indexLeaves bool
	validate    bool
}

type ConfigBuilder struct {
	config *Config
}

func NewConfig() *Config {

	config := &Config{
		indexLeaves: _defIndexLeaves,
		validate:    _defValidate,
	}

	return config
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{config: NewConfig()}
}

func (c *ConfigBuilder) SetIndexLeaves(value bool) {
	c.config.indexLeaves = value
}

func (c *ConfigBuilder) SetValidate(value bool) {
	c.config.validate = true
}

func (c *ConfigBuilder) Build() *Config {
	return c.config
}
