package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"mirasynth.stream/notifier/internal/atlas"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "notifier [command]",
		Short: atlas.NOTIFIER_SHORT_DESC,
		Long:  atlas.NOTIFIER_LONG_DESC,
	}

	rootCmd.AddCommand(NewPushCmd())
	rootCmd.AddCommand(NewServerhCmd())
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
