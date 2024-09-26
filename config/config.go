package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	GinMode           string   `mapstructure:"GIN_MODE"`
	DBSource          string   `mapstructure:"DB_SOURCE"`
	MigrationURL      string   `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress string   `mapstructure:"HTTP_SERVER_ADDRESS"`
	Domain            string   `mapstructure:"DOMAIN"`
	AllowedOrigins    []string `mapstructure:"ALLOWED_ORIGINS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
