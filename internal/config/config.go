package config

import (
	"bytes"
	"encoding/json"

	"github.com/spf13/viper"
)

type config struct {
	DiscordWebhook string `json:"discordWebhook"`
}

func SetupConfig() {
	defaultConfig := &config{
		DiscordWebhook: "",
	}

	defaultConfigBytes, err := json.Marshal(&defaultConfig)
	if err != nil {
		panic(err)
	}

	viper.ReadConfig(bytes.NewReader(defaultConfigBytes))

	viper.SetConfigType("json")
	viper.SetEnvPrefix("notifier")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

// NOTIFIER_DISCORDWEBHOOK
func GetDiscordWebhook() string {
	return viper.GetString("discordWebhook")
}
