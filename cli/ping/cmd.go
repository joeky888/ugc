package ping

import (
	"github.com/joeky888/ugc/tool"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:                "ping",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		tool.CaptureWorker(NewConfig())
	},
}
