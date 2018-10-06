package webhooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/muneebabbas/homebot-core/config"
	"github.com/muneebabbas/homebot-core/utils"
)

var (
	// Discord The discordgo Session object
	Discord *discordgo.Session
	// TestChannel Discord channel id for testing purposes
	TestChannel = "498093617185816626"
	// LogsChannel Discord channel for different logs like updates etc.
	LogsChannel = "493441003144085512"
	// MediaChannel Discord channel for updates about media libraries
	MediaChannel = "493440929374404618"
	// err Error
	err error
)

func init() {
	Discord, err = discordgo.New("Bot " + config.BotToken)
	utils.HandleError("Error creating discord engine", err)

	err = Discord.Open()
	utils.HandleError("Error connecting to discord gateway", err)
}
