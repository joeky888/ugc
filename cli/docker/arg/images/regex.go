package images

import (
	"regexp"

	"github.com/joeky888/ugc/tool"
)

// NewConfig df regex config
func NewConfig() []tool.Conf {
	return []tool.Conf{
		{
			// REPO, TAG, IMAGE ID
			Regex:  regexp.MustCompile(`^([a-z]+\/?[^\s]+)(\s+)([^\s]+)(\s+)(\w+)`),
			Colors: []string{tool.Default, tool.Default, tool.Default, tool.Cyan, tool.Gray},
		},
	}
}
