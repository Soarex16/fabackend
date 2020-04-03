package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config - global application configuration
type Config struct {
	Port               int
	DbConnectionString string
	JwtSecret          []byte
}

// InitConfiguration - reads config and performs validation
func InitConfiguration() (*Config, error) {
	cfg := &Config{
		Port:               viper.GetInt("port"),
		DbConnectionString: viper.GetString("dbConnectionString"),
		JwtSecret:          []byte(viper.GetString("jwtSecret")),
	}

	if cfg.Port == 0 {
		return nil, fmt.Errorf("'port' must be specified in configuration file")
	}

	if len(cfg.DbConnectionString) == 0 {
		return nil, fmt.Errorf("'dbConnectionString' must be specified in configuration file")
	}

	if len(cfg.JwtSecret) == 0 {
		return nil, fmt.Errorf("'jwtSecret' must be specified in configuration file")
	}

	return cfg, nil
}
