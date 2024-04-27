package discord

import (
	"github.com/spf13/cobra"
	"mirasynth.stream/notifier/internal/atlas"
	"mirasynth.stream/notifier/internal/push"
	"mirasynth.stream/notifier/internal/server/notify/services/common"
	"mirasynth.stream/notifier/internal/server/notify/services/discord"
)

func NewPushDiscordCmd() *cobra.Command {
	pushData := push.PushData{}

	pushCmd := &cobra.Command{
		Use:   "discord -c context [text]",
		Short: atlas.PUSH_DISCORD_COMMAND_SHORT_DESC,
		Long:  atlas.PUSH_DISCORD_COMMAND_LONG_DESC,
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			pushData.Text = args[0]

			dns := discord.NewNotifyService()
			err := dns.Notify(&common.Notify{
				Context: pushData.Context,
				Content: pushData.Text,
			})

			return err
		},
	}

	pushCmd.Flags().StringVarP(&pushData.Context, "context", "c", "notifier", atlas.PUSH_DISCORD_COMMAND_FLAG_CONTEXT_DESC)

	return pushCmd
}
