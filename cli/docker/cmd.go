package docker

import (
	"github.com/joeky888/ugc/cli/docker/arg/images"
	"github.com/joeky888/ugc/cli/docker/arg/ps"
	"github.com/joeky888/ugc/tool"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:                "docker",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		tool.CaptureWorker(NewConfig())
	},
}

func init() {
	Cmd.AddCommand(ps.Cmd)
	Cmd.AddCommand(images.Cmd)
}
