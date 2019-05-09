package webhooks

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/muneebabbas/homebot-core/config"
)

// scriptWebhook Handle /webhooks/script endpoint
// This endpoint receives updates from different maintenance scripts like backups
func scriptWebHook(c *gin.Context) {
	discordChannel := LogsChannel

	// Check that all required parameters are present
	var json ScriptJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Switch to test Channel if debug mode or Debug is true in request data
	if config.GinMode == "debug" || json.Debug == true {
		discordChannel = TestChannel
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
	// Check that all required parameters are present
	var json SonarrJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Switch to test Channel if debug mode or Debug is true in request data
	if config.GinMode == "debug" || json.Debug == true {
		discordChannel = TestChannel
	}

	var messageEmbed *discordgo.MessageEmbed
	if json.EventType == "Grab" {
		messageEmbed = createSonarrGrabEmbed(json)
	} else if json.EventType == "Download" {
		messageEmbed = createSonarrDownloadEmbed(json)
	} else if json.EventType == "Test" {
		messageEmbed = NewEmbed().SetTitle("Sonarr Webhook Test").MessageEmbed
		discordChannel = TestChannel
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

// radarrWebhook Handle /webhooks/radarr
// This function handles updates from radarr
func radarrWebhook(c *gin.Context) {
	discordChannel := MediaChannel
	// Check that all required parameters are present
	var json RadarrJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Switch to test Channel if debug mode or Debug is true in request data
	if config.GinMode == "debug" || json.Debug == true {
		discordChannel = TestChannel
	}
	fmt.Println(json.Debug)

	var messageEmbed *discordgo.MessageEmbed
	if json.EventType == "Grab" {
		messageEmbed = createRadarrGrabEmbed(json)
	} else if json.EventType == "Download" {
		messageEmbed = createRadarrDownloadEmbed(json)
	} else if json.EventType == "Test" {
		messageEmbed = NewEmbed().SetTitle("Radarr Webhook Test").MessageEmbed
		discordChannel = TestChannel
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

// embedTestWebhook Test webhook for checking Discord's Embed formatting
func embedTestWebhook(c *gin.Context) {
	messageEmbed := createTestEmbed()
	_, err := Discord.ChannelMessageSendEmbed(TestChannel, messageEmbed)
	if err != nil {
		fmt.Println("error sending message to discord channel,", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error while sending message to discord"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
