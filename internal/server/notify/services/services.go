package services

import (
	"mirasynth.stream/notifier/internal/server/notify/services/common"
	"mirasynth.stream/notifier/internal/server/notify/services/discord"
	"mirasynth.stream/notifier/internal/server/notify/services/system"
)

var notifyServices = make(map[string]common.NotifyService)

func StartServices() {
	notifyServices[discord.SERVICE_NAME] = discord.NewNotifyService()
	notifyServices[system.SERVICE_NAME] = system.NewNotifyService()
}

func NotifyServices(notify *common.Notify) {
	notifyServices[discord.SERVICE_NAME].Notify(notify)
	notifyServices[system.SERVICE_NAME].Notify(notify)
}
