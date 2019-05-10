package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/muneebabbas/homebot-core/config"
	"github.com/muneebabbas/homebot-core/webhooks"
)

// bodyLogger Gin middleware to log the body of a request
// Currently only works with json
func bodyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			var prettyJSON bytes.Buffer
			json.Indent(&prettyJSON, bodyBytes, "", "    ")
			fmt.Fprintln(gin.DefaultWriter, string(prettyJSON.Bytes()))
		}

		// Point Body to new reader as the previous reader was consumed
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	}
}

func main() {
	gin.SetMode(config.GinMode)
	gin.ForceConsoleColor()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(bodyLogger())
	router.Use(gin.Logger())
	// Register webhook group routes
	webhooksGroup := router.Group("/webhooks")
	webhooks.RegisterWebhookRoutes(webhooksGroup)

	router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
}
