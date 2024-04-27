package cmd

import (
	"github.com/spf13/cobra"
	"mirasynth.stream/notifier/cmd/push/discord"
	"mirasynth.stream/notifier/internal/atlas"
	"mirasynth.stream/notifier/internal/push"
)

func NewPushCmd() *cobra.Command {
	pushData := push.PushData{}

	pushCmd := &cobra.Command{
		Use:   "push -c context -i icon [title] [text]",
		Short: atlas.PUSH_COMMAND_SHORT_DESC,
		Long:  atlas.PUSH_COMMAND_LONG_DESC,
		Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			pushData.Title = args[0]
			pushData.Text = args[1]

			x := push.NewPushService()

			if pushData.Critical {
				return x.PushCritical(&pushData)
			}

			return x.Push(&pushData)
		},
	}

	pushCmd.Flags().StringVarP(&pushData.Context, "context", "c", "", atlas.PUSH_COMMAND_FLAG_CONTEXT_DESC)
	pushCmd.Flags().StringVarP(&pushData.Icon, "icon", "i", "", atlas.PUSH_COMMAND_FLAG_ICON_DESC)
	pushCmd.Flags().BoolVarP(&pushData.Critical, "critical", "u", false, atlas.PUSH_COMMAND_FLAG_CRITICAL_DESC)

	pushCmd.AddCommand(discord.NewPushDiscordCmd())

	return pushCmd
}
