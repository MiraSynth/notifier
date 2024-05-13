package health

import (
	"github.com/gin-gonic/gin"
	"mirasynth.stream/notifier/internal/server/health/live"
	"mirasynth.stream/notifier/internal/server/health/ready"
)

func RegisterController(routerGroup *gin.RouterGroup, rootPath string) {
	healthRouterGroup := routerGroup.Group(rootPath)

	live.RegisterController(healthRouterGroup, "/live")
	ready.RegisterController(healthRouterGroup, "/reaady")
}
