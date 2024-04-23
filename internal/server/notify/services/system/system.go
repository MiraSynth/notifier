package system

import (
	"mirasynth.stream/notifier/internal/push"
	"mirasynth.stream/notifier/internal/server/notify/services/common"
)

const SERVICE_NAME = "system"

type systemNotifyService struct {
}

func NewNotifyService() common.NotifyService {
	var notifyService common.NotifyService = &systemNotifyService{}
	return notifyService
}

func (p *systemNotifyService) Notify(notify *common.Notify) (string, error) {
	pushData := push.PushData{
		Context: notify.Context,
		Title:   "From Webhook",
		Text:    notify.Content,
	}

	x := push.NewPushService()
	_, err := x.Push(&pushData)

	return "", err
}
