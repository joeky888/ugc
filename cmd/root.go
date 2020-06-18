package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Short: "help message",
		Long:  `Usage: ./ugc <COMMAND> <ARGS> <ARGS>...`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
