package webhooks

import "github.com/gin-gonic/gin"

// RegisterWebhookRoutes Register all routes for different webhooks
func RegisterWebhookRoutes(router *gin.RouterGroup) {
	router.POST("/script", scriptWebHook)
	router.POST("/sonarr", sonarrWebhook)
	router.POST("/radarr", radarrWebhook)
	router.POST("/embed-test", embedTestWebhook)
}
