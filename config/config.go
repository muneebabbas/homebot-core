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
	// GinMode The mode of gin
	GinMode string
)

func init() {
	// Initialise viper to read config.yaml
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	utils.HandleError("Error reading config file", err)

	// Read the variables defined by var so that they are available as config.Variable
	BotToken = os.Getenv("BOT_TOKEN")
	Port = viper.GetString("PORT")
	Host = viper.GetString("HOST")
	GinMode = viper.GetString("GIN_MODE")
}
