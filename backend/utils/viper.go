package utils

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Load all configuration needed for the service
func LoadViperConfig() {
	// Load config.yaml
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()

	// Load .env
	_ = godotenv.Load()

	// Bind config to environment variables
	viper.BindEnv("server.host", "SERVER_HOST")
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("algorand.address", "ALGORAND_ADDRESS")
	viper.BindEnv("algorand.api_header", "ALGORAND_API_HEADER")
	viper.BindEnv("algorand.api_token", "ALGORAND_API_TOKEN")

	// Set defaults
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "8081")

	requiredConfigs := []string{
		"algorand.address",
		"algorand.api_header",
		"algorand.api_token",
	}

	var logger = GetLoggerInstance()
	for _, key := range requiredConfigs {
		if viper.Get(key) == nil {
			logger.Warn(fmt.Sprintf("Config key is missing: %s", key))
		}
	}
}
