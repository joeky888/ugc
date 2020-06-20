package cmd

import (
	"github.com/joeky888/ugc/cli/curl"
	"github.com/joeky888/ugc/cli/df"
	"github.com/joeky888/ugc/cli/ping"
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
	// rootCmd.SetHelpCommand(&cobra.Command{
	// 	Use:    "nohelp",
	// 	Hidden: true,
	// })
	rootCmd.AddCommand(curl.Cmd)
	rootCmd.AddCommand(df.Cmd)
	rootCmd.AddCommand(ping.Cmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
