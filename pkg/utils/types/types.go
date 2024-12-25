package types

import "time"

type Config struct {
	DefaultTimeout time.Duration
	CreatePackages bool
}

type CliConfig struct {
	Timeout int16
	Takeover bool
}