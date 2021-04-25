package main

import (
	flashCmd "github.com/jls5177/goqmk/flash/cmd"
	viaCmd "github.com/jls5177/goqmk/via/cmd"
	"github.com/spf13/cobra"
	"log"
)

const AppName = "goqmk"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: AppName + " [flags]",
		Short: "Cli utility to interact with QMK keyboards",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(viaCmd.NewViaCommands())
	cmd.AddCommand(flashCmd.NewFlashCommand())
	return cmd
}

func main() {
	command := NewRootCommand()

	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
