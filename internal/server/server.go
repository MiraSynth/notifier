package server

import (
	"github.com/beego/beego/v2/server/web"
	"mirasynth.stream/notifier/internal/server/notify"
)

func StartServer() {
	web.BConfig.CopyRequestBody = true

	notify.RegisterController("/api/v1/notify")

	web.Run(":3038")
}
