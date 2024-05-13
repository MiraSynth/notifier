package server

import (
	"github.com/gin-gonic/gin"
	"mirasynth.stream/notifier/internal/server/health"
	"mirasynth.stream/notifier/internal/server/notify"
)

func StartServer() {
	ginEngine := gin.Default()

	routerGroup := ginEngine.Group("/api/v1")

	health.RegisterController(routerGroup, "/health")
	notify.RegisterController(routerGroup, "/notify")

	ginEngine.Run(":3038")
}
