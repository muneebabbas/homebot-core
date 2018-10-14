package webhooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/muneebabbas/homebot-core/config"
	"github.com/muneebabbas/homebot-core/utils"
)

var (
	// Discord The discordgo Session object
	Discord *discordgo.Session

	// err Error
	err error
)

const (
	// TestChannel Discord channel id for testing purposes
	TestChannel = "498093617185816626"
	// LogsChannel Discord channel for different logs like updates etc.
	LogsChannel = "493441003144085512"
	// MediaChannel Discord channel for updates about media libraries
	MediaChannel = "493440929374404618"

	// Discord colors
	successColor = 0x4CAF50
	errorColor   = 0xDD2C00
	infoColor    = 0x4286f4

	// Logos
	sonarrLogo = "https://i.imgur.com/7JooOj9.png"
	radarrLogo = "https://i.imgur.com/iR79RP1.png"
	backupLogo = "https://i.imgur.com/cpz7Hj9.png"

	// Other constants
	logLines  = 3
	checkMark = "\u2713"
	crossMark = "\u2717"
)

func init() {
	Discord, err = discordgo.New("Bot " + config.BotToken)
	utils.HandleError("Error creating discord engine", err)

	err = Discord.Open()
	utils.HandleError("Error connecting to discord gateway", err)
}
