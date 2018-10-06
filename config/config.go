package config

import (
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
)

func init() {
	// Initialise viper to read config.yaml
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	utils.HandleError("Error reading config file", err)
	BotToken = viper.GetString("BOT_TOKEN")
	Port = viper.GetString("PORT")
	Host = viper.GetString("HOST")
}
