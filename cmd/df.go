package cmd

import (
	"github.com/joeky888/ugc/color/df"
	"github.com/joeky888/ugc/tool"
	"github.com/spf13/cobra"
)

var dfCmd = &cobra.Command{
	Use: "df",
	ValidArgs: []string{
		"-h\tanother",
		"-a\tdescription for rs",
		"-b",
	},
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		tool.CaptureWorker(df.NewConfig())
	},
}

func init() {
	rootCmd.AddCommand(dfCmd)
}
