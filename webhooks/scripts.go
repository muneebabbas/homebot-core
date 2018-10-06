package webhooks

import (
	"fmt"
	"net/http"

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
}

func scriptWebHook(c *gin.Context) {
	var json ScriptJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message := fmt.Sprintf("The script %s was ran successfully", json.Path)
	_, err := Discord.ChannelMessageSend(TestChannel, message)
	if err != nil {
		fmt.Println("error sending message to discord channel,", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error while sending message to discord"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
	return
}
