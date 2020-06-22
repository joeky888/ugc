package ps

import (
	"github.com/joeky888/ugc/tool"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:                "ps",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		tool.CaptureWorker(NewConfig())
	},
}
