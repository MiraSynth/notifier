package health

import (
	"github.com/gin-gonic/gin"
	"mirasynth.stream/notifier/internal/server/health/live"
	"mirasynth.stream/notifier/internal/server/health/ready"
)

func RegisterController(ge *gin.Engine) {
	live.RegisterController(ge, "/api/v1/health/live")
	ready.RegisterController(ge, "/api/v1/health/reaady")
}
