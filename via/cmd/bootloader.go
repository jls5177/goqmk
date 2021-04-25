package cmd

import (
	"github.com/jls5177/goqmk/via/cmd/flags"
	"github.com/spf13/cobra"
)

type BootloaderOptions struct {
	flags.UsbDevice
}

func NewBootloaderCommand() * cobra.Command {
	opts := BootloaderOptions{}

	cmd := &cobra.Command{
		Use: "bootloader [flags]",
		Aliases: []string{ "bl" },
		Short: "Resets the QMk keyboard into the bootloader",
		Args: cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.UsbDevice.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.UsbDevice.Device().GotoBootloader()
		},
	}

	opts.UsbDevice.AddFlags(cmd)

	return cmd
}
