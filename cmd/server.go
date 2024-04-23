package cmd

import (
	"github.com/spf13/cobra"
	"mirasynth.stream/notifier/internal/atlas"
	"mirasynth.stream/notifier/internal/server"
)

func NewServerhCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: atlas.SERVER_COMMAND_SHORT_DESC,
		Long:  atlas.SERVER_COMMAND_LONG_DESC,
		Run: func(cmd *cobra.Command, args []string) {
			server.StartServer()
		},
	}

	return cmd
}
