package services

import (
	"mirasynth.stream/notifier/internal/server/notify/services/common"
	"mirasynth.stream/notifier/internal/server/notify/services/discord"
)

var notifyServices = make(map[string]common.NotifyService)

func StartServices() {
	notifyServices[discord.SERVICE_NAME] = discord.NewDiscordNotifyService()
}

func NotifyServices(notify *common.Notify) {
	notifyServices[discord.SERVICE_NAME].Notify(notify)
}
