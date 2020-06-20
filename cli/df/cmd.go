package df

import (
	"github.com/joeky888/ugc/tool"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:                "df",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		tool.CaptureWorker(NewConfig())
	},
}

func init() {
	Cmd.Flags().BoolP("portability", "P", false, "POSIX output format")
	Cmd.Flags().BoolP("block-size", "k", false, "1024-byte blocks (default)")
}
