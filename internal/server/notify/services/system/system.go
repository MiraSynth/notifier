package system

import (
	"mirasynth.stream/notifier/internal/push"
	"mirasynth.stream/notifier/internal/server/notify/services/common"
)

const SERVICE_NAME = "system"

type systemNotifyService struct {
	pushService *push.Push
}

func NewNotifyService() common.NotifyService {
	pushService := push.NewPushService()
	var notifyService common.NotifyService = &systemNotifyService{
		pushService: &pushService,
	}
	return notifyService
}

func (p *systemNotifyService) Notify(notify *common.Notify) error {
	pushData := push.PushData{
		Context: notify.Context,
		Title:   "From Webhook",
		Text:    notify.Content,
	}

	pushService := *p.pushService
	err := pushService.Push(&pushData)

	return err
}
