package lighting

import (
	"github.com/spf13/cobra"
	"github.com/jls5177/goqmk/via/cmd/flags"
	"strconv"
)

type BrightnessOptions struct {
	brightnessValue int
	flags.UsbDevice
}

func NewBrightnessCommand() *cobra.Command {
	opts := BrightnessOptions{}

	cmd := &cobra.Command{
		Use: "brightness [flags] VALUE",
		Aliases: []string{ "br", "b" },
		Short: "Sets the current RGB lighting brightness",
		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			value, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			opts.brightnessValue = value
			return opts.UsbDevice.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.UsbDevice.Device().Lighting().Brightness(uint8(opts.brightnessValue))
		},
	}

	opts.UsbDevice.AddFlags(cmd)

	return cmd
}
