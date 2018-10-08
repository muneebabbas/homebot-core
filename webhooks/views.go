package webhooks

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

// ScriptJSON The structure of the json the script webhook expects
type ScriptJSON struct {
	Path        string `json:"path" binding:"required"`
	Description string `json:"description" binding:"required"`
	Time        *int   `json:"time" binding:"exists"`
	Returncode  *int   `json:"returncode" binding:"exists"`
	Stderr      string `json:"stderr" binding:"required"`
	Stdout      string `json:"stdout" binding:"required"`
	Logfile     string `json:"logfile" binding:"required"`
}

// scriptWebhook Handle /webhooks/script endpoint
// This endpoint receives updates from different maintenance scripts like backups
func scriptWebHook(c *gin.Context) {
	var json ScriptJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	embed := createScriptEmbed(json)
	_, err := Discord.ChannelMessageSendEmbed(TestChannel, embed)
	if err != nil {
		fmt.Println("error sending message to discord channel,", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error while sending message to discord"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
	return
}

// createScriptEmbed Use Embed struct and it's functions to create a MessageEmbed for scripts webhook
func createScriptEmbed(json ScriptJSON) *discordgo.MessageEmbed {
	embed := NewEmbed().SetTimestamp()
	// Set title and color  based on Returncode
	if *json.Returncode == 0 {
		embed.SetTitle("Success: Executed successfully").SetColor(successColor)
	} else {
		title := fmt.Sprintf("Error: Exited with returncode %d", *json.Returncode)
		embed.SetTitle(title).SetColor(errorColor)
	}

	// Add fields for script path and execution time
	embed.AddField("Path", "Script is located at "+json.Path).
		AddField("Execution Time", fmt.Sprintf("Script took %d seconds to execute", *json.Time)).
		SetFooter(fmt.Sprintf("Please see the logs at %s for more details", json.Logfile))
	return embed.MessageEmbed
}

// sonarrWebhook Handle /webhooks/sonarr
// This function handles updates from sonarr
// func sonarrWebhook(c *gin.Context) {

// }
