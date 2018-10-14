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
	Debug       bool   `json:"debug" binding:"-"`
	Path        string `json:"path" binding:"required"`
	Description string `json:"description" binding:"required"`
	Time        *int   `json:"time" binding:"exists"`
	Returncode  *int   `json:"returncode" binding:"exists"`
	Stderr      string `json:"stderr" binding:"-"`
	Stdout      string `json:"stdout" binding:"-"`
	Logfile     string `json:"logfile" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

// SonarrJSON The structure of Sonarr's webhook requests
type SonarrJSON struct {
	Debug     bool   `json:"debug" binding:"-"`
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

// RadarrJSON The structure of Radarr's webhook requests
type RadarrJSON struct {
	Debug     bool   `json:"debug" binding:"-"`
	EventType string `json:"eventType" binding:"required"`
	Movie     struct {
		Title string `json:"title" binding:"required"`
	}
	Release struct {
		Quality       string `json:"quality" binding:"required"`
		QuaityVersion *int   `json:"qualityVersion" binding:"required"`
		Size          *int   `json:"size" binding:"required"`
		ReleaseTitle  string `json:"releaseTitle" binding:"required"`
		ReleaseGroup  string `json:"releaseGroup" binding:"required"`
		Indexer       string `json:"indexer" binding:"required"`
	} `json:"release" binding:"-"`

	MovieFile struct {
		RelativePath string `json:"relativePath" binding:"required"`
		Path         string `json:"path" binding:"required"`
		Quality      string `json:"quality" binding:"required"`
		ReleaseGroup string `json:"releaseGroup" binding:"required"`
	} `json:"movieFile" binding:"-"`

	RemoteMovie struct {
		Title string `json:"title" binding:"required"`
		Year  *int   `json:"year" binding:"required"`
	} `json:"remoteMovie" binding:"required"`
}

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
