package lighting

import (
	"github.com/spf13/cobra"
	"github.com/jls5177/goqmk/via/cmd/flags"
)

type SaveOptions struct {
	colorValue int
	flags.UsbDevice
}

func NewSaveCommand() *cobra.Command {
	opts := SaveOptions{}

	cmd := &cobra.Command{
		Use: "save [flags]",
		Aliases: []string{ "sa" },
		Short: "Saves the current RGB lighting state to EEPROM",
		Args: cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.UsbDevice.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.UsbDevice.Device().Lighting().Save()
		},
	}

	opts.UsbDevice.AddFlags(cmd)

	return cmd
}
