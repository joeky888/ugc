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
			// HTTP header
			Regex: regexp.MustCompile(`(?m)^([A-Z].*)\:\s(.*)$`),
			Color: tool.Blue,
		},
		{
			// HTTP success
			Regex: regexp.MustCompile(`(?m)^HTTP.*?\s([1-3]\d{2}.*)$`),
			Color: tool.Green,
		},
		{
			// HTTP error
			Regex: regexp.MustCompile(`(?m)^HTTP.*?\s([4|5]\d{2}.*)$`),
			Color: tool.Red,
		},
		{
			// JSON attribute
			Regex: regexp.MustCompile(`(\".*\")\:`),
			Color: tool.Blue,
		},
		// {
		// 	// String value
		// 	Regex: regexp.MustCompile(`\s*(\".*\")`),
		// 	Color: tool.Yellow,
		// },
		// {
		// 	// Number value
		// 	Regex: regexp.MustCompile(`\s*(\d+)`),
		// 	Color: tool.Purple,
		// },
		{
			// Boolean value
			Regex: regexp.MustCompile(`\s*(true|false)`),
			Color: tool.Purple,
		},
		{
			// Null value
			Regex: regexp.MustCompile(`\s*(null)`),
			Color: tool.Red,
		},
	}
}
