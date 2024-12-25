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

type AddAccessTokenOptions struct {
	Email	string
	Token	string
}

type AccountCreationOptions struct {
	Email			string
	Username		string
	Password		string
	Registry		string
}