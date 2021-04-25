package lighting

import (
	"github.com/spf13/cobra"
	"github.com/jls5177/goqmk/via/cmd/flags"
	"strconv"
)

type EffectOptions struct {
	effectValue int
	flags.UsbDevice
}

func NewEffectCommand() *cobra.Command {
	opts := EffectOptions{}

	cmd := &cobra.Command{
		Use: "effect [flags] VALUE",
		Aliases: []string{ "e" },
		Short: "Sets the current RGB lighting effect",
		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			value, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			opts.effectValue = value
			return opts.UsbDevice.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.UsbDevice.Device().Lighting().Effect(uint8(opts.effectValue))
		},
	}

	opts.UsbDevice.AddFlags(cmd)

	return cmd
}
