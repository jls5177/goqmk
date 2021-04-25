package lighting

import (
	"github.com/spf13/cobra"
)

func NewLightingCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "lighting",
		Aliases: []string{ "lt", "l" },
		Short: "Commands that interact with RGB lights in QMK keyboards",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(NewBrightnessCommand())
	cmd.AddCommand(NewColorCommand())
	cmd.AddCommand(NewEffectCommand())
	cmd.AddCommand(NewSpeedCommand())
	cmd.AddCommand(NewSaveCommand())

	return cmd
}
