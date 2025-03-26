package config

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	ENV_MODE_DEVELOPMENT = iota + 1
	ENV_MODE_PRODUCTION  = iota + 1
	ENV_MODE_STAGE       = iota + 1
)

const (
	envModeDevelopmentStr = "development"
	envModeProductionStr  = "production"
	envModeStageStr       = "stage"
)

type Config struct {
	EnvMode uint8
	Port    uint16
	Host    string
	DB      DataBaseConfig
	Client  ClientConfig
}

type DataBaseConfig struct {
	Host     string
	Port     uint16
	Name     string
	User     string
	Password string
	SSLMode  string
}

type ClientConfig struct {
	RulesURL    string
	ProjectsURL string
	ParserURL   string
}

func LoadConfig(envMode string) (*Config, error) {
	mode, err := validateEnvMode(envMode)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	// Set environment variables
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("PORT", 8787)
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_USER", "user")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_SSLMODE", "disable")

	// Map environment variables to config structure
	config.Port = uint16(viper.GetInt("PORT"))
	config.Host = viper.GetString("HOST")
	config.DB = DataBaseConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     uint16(viper.GetInt("DB_PORT")),
		Name:     viper.GetString("DB_NAME"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		SSLMode:  viper.GetString("DB_SSLMODE"),
	}

	config.EnvMode = mode

	return config, nil
}

func MustLoadConfig(envMode string) *Config {
	config, err := LoadConfig(envMode)
	if err != nil {
		panic(err)
	}

	return config
}

func validateEnvMode(envMode string) (uint8, error) {
	var mode uint8
	switch envMode {
	case envModeDevelopmentStr:
		mode = ENV_MODE_DEVELOPMENT
	case envModeProductionStr:
		mode = ENV_MODE_PRODUCTION
	case envModeStageStr:
		mode = ENV_MODE_STAGE
	default:
		return mode, errors.New("Unknown environment mode")
	}

	return mode, nil
}
