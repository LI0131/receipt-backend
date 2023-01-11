package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Environment    string
	PublicPort     string
	LogLevel       string
	Hostname       string
	BaseApiPath    string
	DatabaseConfig DbConfig
}

type DbConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func Get() *Config {
	options := viper.New()

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "0.0.0.0"
	}

	// Set Basic Defaults
	options.SetDefault("Environment", "Dev")
	options.SetDefault("Hostname", hostname)
	options.SetDefault("PublicPort", 8080)
	options.SetDefault("LogLevel", "INFO")
	options.SetDefault("BaseApiPath", "/v1/api")

	// Set Database Defaults
	options.SetDefault("DatabaseUser", "receipt-api")
	options.SetDefault("DatabasePassword", "receipt-api")
	options.SetDefault("DatabaseName", "receipt-api")
	options.SetDefault("DatabaseHost", "0.0.0.0")
	options.SetDefault("DatabasePort", 5432)

	return &Config{
		Environment: options.GetString("Environment"),
		PublicPort:  options.GetString("PublicPort"),
		LogLevel:    options.GetString("LogLevel"),
		Hostname:    options.GetString("Hostname"),
		BaseApiPath: options.GetString("BaseApiPath"),
		DatabaseConfig: DbConfig{
			DBUser:     options.GetString("DatabaseUser"),
			DBPassword: options.GetString("DatabasePassword"),
			DBName:     options.GetString("DatabaseName"),
			DBHost:     options.GetString("DatabaseHost"),
			DBPort:     options.GetString("DatabasePort"),
		},
	}
}
