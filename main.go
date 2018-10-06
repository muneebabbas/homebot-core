package main

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	botToken string
	port     string
	host     string

	// TestChannel Discord channel id for testing purposes
	TestChannel = "498093617185816626"
	// LogsChannel Discord channel for different logs like updates etc.
	LogsChannel = "493441003144085512"
	// MediaChannel Discord channel for updates about media libraries
	MediaChannel = "493440929374404618"
)

// ScriptJSON The structure of the json the script webhook expects
type ScriptJSON struct {
	Path        string `json:"path" binding:"required"`
	Description string `json:"description" binding:"required"`
	Time        *int   `json:"time" binding:"exists"`
	Returncode  *int   `json:"returncode" binding:"exists"`
	Stderr      string `json:"stderr" binding:"required"`
	Stdout      string `json:"stdout" binding:"required"`
}

func init() {
	// Initialise viper to read config.yaml
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	handleError("Error reading config file", err)
	botToken = viper.GetString("BOT_TOKEN")
	port = viper.GetString("PORT")
	host = viper.GetString("HOST")
}

func main() {
	router := gin.Default()
	discord, err := discordgo.New("Bot " + botToken)
	handleError("Error creating discord engine", err)

	err = discord.Open()
	handleError("Error connecting to discord gateway", err)

	// Webhooks for handling requests from scripts, sonarr and radarr etc.
	webhooks := router.Group("/webhooks")
	webhooks.POST("/script", func(c *gin.Context) {
		var json ScriptJSON
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		message := fmt.Sprintf("The script %s was ran successfully", json.Path)
		_, err := discord.ChannelMessageSend(TestChannel, message)
		if err != nil {
			fmt.Println("error sending message to discord channel,", err)
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "Error while sending message to discord"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
		return
	})

	router.Run(host + ":" + port)
}

func handleError(message string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", message, err)
		panic(err)
	}
}
