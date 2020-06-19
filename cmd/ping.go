package cmd

import (
	"github.com/joeky888/ugc/color/ping"
	"github.com/joeky888/ugc/tool"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:                "ping",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		tool.CaptureWorker(ping.NewConfig())
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
