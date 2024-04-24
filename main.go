package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"mirasynth.stream/notifier/cmd"
	"mirasynth.stream/notifier/internal/config"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	config.SetupConfig()
	cmd.Execute()
}
