package server

import (
	"github.com/gin-gonic/gin"
	"mirasynth.stream/notifier/internal/server/health"
	"mirasynth.stream/notifier/internal/server/notify"
)

func StartServer() {
	ge := gin.Default()

	health.RegisterController(ge)
	notify.RegisterController(ge, "/api/v1/notify")

	ge.Run(":3038")
}
