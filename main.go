package main

import (
	"mirasynth.stream/notifier/cmd"
	"mirasynth.stream/notifier/internal/config"
)

func main() {
	config.SetupConfig()
	cmd.Execute()
}
