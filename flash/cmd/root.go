package cmd

import (
	"errors"
	"fmt"
	"github.com/jls5177/goqmk/flash"
	"github.com/jls5177/goqmk/via/cmd/flags"
	"github.com/spf13/cobra"
	"os"
)

type FlashOptions struct {
	firmwarePath string
	flags.UsbDevice
}

func NewFlashCommand() *cobra.Command {
	opts := FlashOptions{}

	cmd := &cobra.Command{
		Use: "flash [flags] FIRMWARE_FILE",
		Short: "Flashes the given flash file to the QMK keyboard",
		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.firmwarePath = args[0]
			if err := validateFilePath(opts.firmwarePath); err != nil {
				return err
			}
			return opts.UsbDevice.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return flash.FlashBoard(opts.firmwarePath, opts.Device())
		},
	}

	opts.UsbDevice.AddFlags(cmd)

	return cmd
}

func validateFilePath(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("the file path you specified does not exist: %v", path)
	} else if err != nil {
		return fmt.Errorf("error parsing path: %w", err)
	}
	return nil
}