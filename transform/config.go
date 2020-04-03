package transform

const (
	_defIndexLeaves = true
)

type Config struct {
	indexLeaves bool
}

type ConfigBuilder struct {
	config *Config
}

func NewConfig() *Config {

	config := &Config{
		indexLeaves: _defIndexLeaves,
	}

	return config
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{config: NewConfig()}
}

func (c *ConfigBuilder) SetIndexLeaves(value bool) {
	c.config.indexLeaves = value
}

func (c *ConfigBuilder) Build() *Config {
	return c.config
}
