package scan

import "time"

type Config struct {
	DefaultTimeout time.Duration
	CreatePackages bool
	Monitoring 	bool
}

type CliConfig struct {
	Timeout int16
	Takeover bool
}

func CreateScanConfig(config CliConfig) Config {
	return Config{
		DefaultTimeout: time.Duration(config.Timeout) * time.Millisecond,
		CreatePackages: config.Takeover,
		Monitoring: false,
	}
}