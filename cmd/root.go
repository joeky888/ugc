package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ugc",
		Short: "help message",
		Long:  `Usage: ./ugc <COMMAND> <ARGS> <ARGS>...`,
	}
)

func init() {
	// Disable help command
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "nohelp",
		Hidden: true,
	})
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
