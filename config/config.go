package config

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./go_dig_sample.db",
		Port:         "8000",
	}
}
