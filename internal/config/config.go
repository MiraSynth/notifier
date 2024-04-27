package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var venv *viper.Viper

func SetupConfig() {
	venv = viper.New()
	venv.SetEnvPrefix("notifier")
	venv.AutomaticEnv()
}

// NOTIFIER_DISCORD_[CUSTOM_NAME]
func GetDiscordWebhook(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ToLower(name)
	return venv.GetString(fmt.Sprintf("discord_%s", name))
}
