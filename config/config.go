package config

import (
	"os"
)

// Config for database
type Config struct {
	Username     string
	Password     string
	DatabaseName string
	URL          string
}

// GetConfig returns configuration from environment
func GetConfig() *Config {
	return &Config{
		Username:     os.Getenv("DB_USERNAME"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		URL:          os.Getenv("DB_URL"),
	}
}
