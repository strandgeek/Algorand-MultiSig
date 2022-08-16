package viperutil

import (
	"fmt"
	"multisigdb-svc/utils/loggerutil"

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

	viper.BindEnv("auth.jwt_secret", "AUTH_JWT_SECRET")

	viper.BindEnv("logger.level", "LOGGER_LEVEL")
	viper.BindEnv("logger.encoding", "LOGGER_ENCODING")
	viper.BindEnv("logger.output_paths", "LOGGER_OUTPUT_PATHS")
	viper.BindEnv("logger.error_output_paths", "LOGGER_ERROR_OUTPUT_PATHS")

	// Set defaults
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "8081")

	viper.SetDefault("logger.level", "debug")
	viper.SetDefault("logger.encoding", "json")
	viper.SetDefault("logger.output_paths", []string{"stdout", "logs/out.json"})
	viper.SetDefault("logger.error_output_paths", []string{"stdout", "logs/out.json"})

	requiredConfigs := []string{
		"algorand.address",
		"algorand.api_header",
		"algorand.api_token",
		"auth.jwt_secret",
	}

	logger, err := loggerutil.NewLogger()
	if err != nil {
		panic(err)
	}
	for _, key := range requiredConfigs {
		if viper.Get(key) == nil {
			logger.Warn(fmt.Sprintf("Config key is missing: %s", key))
		}
	}
}
