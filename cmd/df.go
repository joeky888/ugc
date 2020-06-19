package cmd

import (
	"github.com/joeky888/ugc/color/df"
	"github.com/joeky888/ugc/tool"
	"github.com/spf13/cobra"
)

var dfCmd = &cobra.Command{
	Use:                "df",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		tool.CaptureWorker(df.NewConfig())
	},
}

func init() {
	dfCmd.Flags().BoolP("portability", "P", false, "POSIX output format")
	dfCmd.Flags().BoolP("block-size", "k", false, "1024-byte blocks (default)")
	rootCmd.AddCommand(dfCmd)
}
