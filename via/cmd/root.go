package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jls5177/goqmk/via/cmd/lighting"
)

func NewViaCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "via",
		Short: "Commands that interact with QMK keyboards using the Via interface",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(NewBootloaderCommand())
	cmd.AddCommand(lighting.NewLightingCommands())

	return cmd
}