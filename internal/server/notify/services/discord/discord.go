package discord

import (
	"bytes"
	"encoding/json"
	"net/http"

	"mirasynth.stream/notifier/internal/server/notify/services/common"
)

const SERVICE_NAME = "discord"

type discordMessage struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type discordNotifyService struct {
}

func NewDiscordNotifyService() common.NotifyService {
	var notifyService common.NotifyService = &discordNotifyService{}
	return notifyService
}

func (p *discordNotifyService) Notify(notify *common.Notify) (string, error) {
	posturl := ""

	dm := discordMessage{
		Username: notify.Context,
		Content:  notify.Content,
	}

	dmRaw, err := json.Marshal(dm)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", posturl, bytes.NewBuffer(dmRaw))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	return "", nil
}
