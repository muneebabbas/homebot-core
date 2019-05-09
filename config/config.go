// Package config is responsible for handling sensitive information like API keys/secrets and configuration varialbes for different deployments
// The spf13/viper package is used to read the config.yaml file from the root directory
package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// Port The Port at which the application runs
	Port string
	// Host The Host at which application runs
	Host string
	// BotToken The Discord Bot Token
	BotToken string
	// GinMode The mode of gin (release vs debug)
	GinMode string
)

func init() {

	// Load environement variables from .env if it's present
	// It's done in a fail safe manner and anything already in the env
	// is preferred
	godotenv.Load()

	// Read environment variables
	BotToken = os.Getenv("BOT_TOKEN")
	Port = os.Getenv("PORT")
	Host = os.Getenv("HOST")
	GinMode = os.Getenv("GIN_MODE")
}
