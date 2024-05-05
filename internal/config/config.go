package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
	"mirasynth.stream/notifier/internal/atlas"
)

var venv *viper.Viper

func SetupConfig() {
	venv = viper.New()
	venv.SetConfigType(atlas.CONFIG_TYPE)
	venv.SetEnvPrefix(atlas.CONFIG_PREFIX)
	venv.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	venv.AutomaticEnv()

	configFilePath, err := verifyConfigFile()
	if err != nil {
		panic(err)
	}

	venv.SetConfigFile(configFilePath)

	err = venv.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("usign config", venv.ConfigFileUsed())
}

func verifyConfigFile() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configDirPath := path.Join(userConfigDir, atlas.CONFIG_NAMESPACE, atlas.CONFIG_PREFIX)
	configFilePath := path.Join(configDirPath, fmt.Sprintf("%s.%s", atlas.CONFIG_FILENAME, atlas.CONFIG_TYPE))

	err = os.MkdirAll(configDirPath, 0755)
	if err != nil {
		return "", err
	}

	_, err = os.Stat(configFilePath)
	if err == nil || !errors.Is(err, os.ErrNotExist) {
		return configFilePath, err
	}

	bytes := []byte{}
	err = os.WriteFile(configFilePath, bytes, 0755)
	if err != nil {
		panic(err)
	}

	return configFilePath, nil
}

// NOTIFIER_DISCORD_[CUSTOM_NAME]
func GetDiscordWebhook(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ToLower(name)
	return venv.GetString(fmt.Sprintf("discord_%s", name))
}
