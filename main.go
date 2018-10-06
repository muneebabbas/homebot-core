package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/muneebabbas/homebot-core/config"
	"github.com/muneebabbas/homebot-core/webhooks"
)

func main() {
	router := gin.Default()

	// Register webhook group routes
	webhooksGroup := router.Group("/webhooks")
	webhooks.RegisterWebhookRoutes(webhooksGroup)

	router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
}
