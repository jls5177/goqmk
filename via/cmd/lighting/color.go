package lighting

import (
	"github.com/spf13/cobra"
	"github.com/jls5177/goqmk/via/cmd/flags"
	"strconv"
)

type ColorOptions struct {
	hue int
	saturation int
	flags.UsbDevice
}

func NewColorCommand() * cobra.Command {
	opts := ColorOptions{}

	cmd := &cobra.Command{
		Use: "color [flags] HUE SATURATION",
		Aliases: []string{ "c" },
		Short: "Sets the RGB color value",
		Args: cobra.ExactArgs(2),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			value, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			opts.hue = value

			value, err = strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			opts.saturation = value

			return opts.UsbDevice.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.UsbDevice.Device().Lighting().Color(uint8(opts.hue), uint8(opts.saturation))
		},
	}

	opts.UsbDevice.AddFlags(cmd)

	return cmd
}
