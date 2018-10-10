package webhooks

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"

	"github.com/gin-gonic/gin"
	"github.com/muneebabbas/homebot-core/config"
)

// ScriptJSON The structure of json the script webhook expects
type ScriptJSON struct {
	Path        string `json:"path" binding:"required"`
	Description string `json:"description" binding:"required"`
	Time        *int   `json:"time" binding:"exists"`
	Returncode  *int   `json:"returncode" binding:"exists"`
	Stderr      string `json:"stderr" binding:"-"`
	Stdout      string `json:"stdout" binding:"-"`
	Logfile     string `json:"logfile" binding:"required"`
}

// SonarrJSON The structure of Sonarr's webhook requests
type SonarrJSON struct {
	EventType string `json:"eventType" binding:"required"`
	Series    struct {
		Title string `json:"title" binding:"required"`
	} `json:"series" binding:"required"`
	Episodes []Episode `json:"episodes" binding:"-"`
	Release  struct {
		Quality       string `json:"quality" binding:"required"`
		QuaityVersion *int   `json:"qualityVersion" binding:"required"`
		Size          *int   `json:"size" binding:"required"`
	} `json:"release" binding:"-"`
	EpisodeFile struct {
		RelativePath string `json:"relativePath" binding:"required"`
		Path         string `json:"path" binding:"required"`
		Quality      string `json:"quality" binding:"required"`
	} `json:"episodeFile" binding:"-"`
	IsUpgrade bool `json:"isUpgrade" binding:"-"`
}

// Episode The structure of a single Episode in Episodes array
type Episode struct {
	Title         string `json:"title" binding:"required"`
	EpisodeNumber *int   `json:"episodeNumber" binding:"required"`
	SeasonNumber  *int   `json:"seasonNumber" binding:"required"`
	Quality       string `json:"quality" binding:"required"`
}

// scriptWebhook Handle /webhooks/script endpoint
// This endpoint receives updates from different maintenance scripts like backups
func scriptWebHook(c *gin.Context) {
	// Switch to test channel for debug mode
	discordChannel := LogsChannel
	if config.GinMode == "debug" {
		discordChannel = TestChannel
	}
	// Check that all required parameters are present
	var json ScriptJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageEmbed := createScriptEmbed(json)
	_, err := Discord.ChannelMessageSendEmbed(discordChannel, messageEmbed)
	if err != nil {
		fmt.Println("error sending message to discord channel,", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error while sending message to discord"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
	return
}

// sonarrWebhook Handle /webhooks/sonarr
// This function handles updates from sonarr
func sonarrWebhook(c *gin.Context) {
	discordChannel := MediaChannel
	if config.GinMode == "debug" {
		discordChannel = TestChannel
	}
	// Check that all required parameters are present
	var json SonarrJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var messageEmbed *discordgo.MessageEmbed
	if json.EventType == "Grab" {
		messageEmbed = createSonarrGrabEmbed(json)
	} else if json.EventType == "Download" {
		messageEmbed = createSonarrDownloadEmbed(json)
	}

	_, err := Discord.ChannelMessageSendEmbed(discordChannel, messageEmbed)
	if err != nil {
		fmt.Println("error sending message to discord channel,", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error while sending message to discord"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
	return

}
