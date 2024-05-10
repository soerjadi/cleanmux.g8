package config

import "time"

var cfg *Config

type Config struct {
	Server   Server
	Database DatabaseConfig
}

type Server struct {
	Port        string
	WaitTimeout int
	BaseURL     string
}

type DatabaseConfig struct {
	Name     string
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	SSL      string
}

func (c Config) WaitTimeout() time.Duration {
	return time.Duration(c.Server.WaitTimeout)
}
