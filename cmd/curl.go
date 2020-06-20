package cmd

import (
	"github.com/joeky888/ugc/color/curl"
	"github.com/joeky888/ugc/tool"
	"github.com/spf13/cobra"
)

var curlCmd = &cobra.Command{
	Use:                "curl",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		tool.CaptureWorker(curl.NewConfig())
	},
}

func init() {
	rootCmd.AddCommand(curlCmd)
}
