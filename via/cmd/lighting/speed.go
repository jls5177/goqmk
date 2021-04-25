package lighting

import (
	"github.com/spf13/cobra"
	"github.com/jls5177/goqmk/via/cmd/flags"
	"strconv"
)

type SpeedOptions struct {
	speedValue int
	flags.UsbDevice
}

func NewSpeedCommand() *cobra.Command {
	opts := SpeedOptions{}

	cmd := &cobra.Command{
		Use: "speed [flags] VALUE",
		Aliases: []string{ "sp" },
		Short: "Sets the current RGB lighting effect speed",
		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			value, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			opts.speedValue = value
			return opts.UsbDevice.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.UsbDevice.Device().Lighting().EffectSpeed(uint8(opts.speedValue))
		},
	}

	opts.UsbDevice.AddFlags(cmd)

	return cmd
}
