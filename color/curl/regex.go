package curl

import (
	"regexp"

	"github.com/joeky888/ugc/tool"
)

// NewConfig curl regex config
// Reference https://michaelheap.com/grc
func NewConfig() []tool.Conf {
	return []tool.Conf{
		{
			// HTTP success
			Regex:  regexp.MustCompile(`(?m)^HTTP.*?\s([1-3]\d{2}.*)$`),
			Colors: []string{tool.Green},
		},
		{
			// HTTP error
			Regex:  regexp.MustCompile(`(?m)^HTTP.*?\s([4|5]\d{2}.*)$`),
			Colors: []string{tool.Red},
		},
		{
			// HTTP header
			Regex:  regexp.MustCompile(`(?m)^([\w-]+)(\:\s*)(.*)`),
			RegexReplace: "$1$2$3",
			Colors: []string{tool.Red, tool.Default, tool.Green},
		},
		{
			// JSON attribute string value
			Regex:        regexp.MustCompile(`(\".*\")(\s*\:\s*)(\".*\")`),
			RegexReplace: "$1$2$3",
			Colors:       []string{tool.Red, tool.Default, tool.Yellow},
		},
		{
			// JSON attribute number/boolean value
			Regex:        regexp.MustCompile(`(\".*\")(\s*\:\s*)(-?[0-9\.]+)`),
			RegexReplace: "$1$2$3",
			Colors:       []string{tool.Red, tool.Default, tool.Blue},
		},
		{
			// JSON attribute boolean value
			Regex:        regexp.MustCompile(`(\".*\")(\s*\:\s*)(true|false)`),
			RegexReplace: "$1$2$3",
			Colors:       []string{tool.Red, tool.Default, tool.Purple},
		},
		{
			// JSON attribute null value
			Regex:        regexp.MustCompile(`(\".*\")(\s*\:\s*)(null)`),
			RegexReplace: "$1$2$3",
			Colors:       []string{tool.Red, tool.Default, tool.Gray},
		},
		{
			// JSON attribute object|array value
			Regex:        regexp.MustCompile(`(\".*\")(\s*\:\s*)(\{|\[)`),
			RegexReplace: "$1$2$3",
			Colors:       []string{tool.Red, tool.Default, tool.Default},
		},
		{
			// Brackets
			Regex: regexp.MustCompile(`\{|\}`),
			Colors: []string{tool.Blue},
		},
	}
}
