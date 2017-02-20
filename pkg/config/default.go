package config

// NewDefaults crates a default config set
func NewDefaults() *Config {
	return &Config{
		Duply: DuplyConfig{Home: "/etd/duply", Root: "/"},
	}
}
