package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/muneebabbas/homebot-core/config"
	"github.com/muneebabbas/homebot-core/utils"
	"github.com/muneebabbas/homebot-core/webhooks"
)

func main() {
	gin.SetMode(config.GinMode)
	gin.ForceConsoleColor()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(utils.BodyLogger())
	router.Use(gin.Logger())
	// Register webhook group routes
	webhooksGroup := router.Group("/webhooks")
	webhooks.RegisterWebhookRoutes(webhooksGroup)

	router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
}
