package cmd

import (
	"fmt"
	"os"

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
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
