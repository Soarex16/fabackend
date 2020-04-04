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

	readEnvOverrides(cfg)

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

// for overriding values from config file via env vars
func readEnvOverrides(cfg *Config) error {
	viper.AutomaticEnv()

	if p := viper.GetInt("LISTEN_PORT"); p != 0 {
		cfg.Port = p
	}

	//try to get connection string from heroku env vars
	if conStr := viper.GetString("DATABASE_URL"); len(conStr) > 0 {
		cfg.DbConnectionString = conStr
	}

	if secret := []byte(viper.GetString("JWT_SECRET")); len(secret) > 0 {
		cfg.JwtSecret = secret
	}
}
