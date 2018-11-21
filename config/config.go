// Package config is responsible for handling sensitive information like API keys/secrets and configuration varialbes for different deployments
// The spf13/viper package is used to read the config.yaml file from the root directory
package config

import (
	"os"

	"github.com/muneebabbas/homebot-core/utils"
	"github.com/spf13/viper"
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
	// Initialise viper to read config.yaml
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	utils.HandleError("Error reading config file", err)

	// BotToken is an environment variable as it shouldn't be visible anywhere
	BotToken = os.Getenv("BOT_TOKEN")

	// Read the variables defined by var so that they are available as config.Variable
	Port = viper.GetString("PORT")
	Host = viper.GetString("HOST")
	GinMode = viper.GetString("GIN_MODE")
}
