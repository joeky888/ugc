package df

import (
	"regexp"

	"github.com/joeky888/ugc/tool"
)

// NewConfig df regex config
func NewConfig() []tool.Conf {
	return []tool.Conf{
		{
			// FS
			Regex:  regexp.MustCompile(`(?m)^(\/[\w]+)+`),
			Colors: []string{tool.Purple},
		},
		{
			// Size '0'
			Regex:  regexp.MustCompile(`\s0\s`),
			Colors: []string{tool.Blue},
		},
		{
			// Size 'K'
			Regex:  regexp.MustCompile(`\s\d*[.,]?\d(K|B)i?\s|\s\d{1,3}\s`),
			Colors: []string{tool.Green},
		},
		{
			// Size 'M'
			Regex:  regexp.MustCompile(`\s\d*[.,]?\dMi?\s|\s\d{4,6}\s`),
			Colors: []string{tool.Yellow},
		},
		{
			// Size 'G'
			Regex:  regexp.MustCompile(`\s\d*[.,]?\dGi?\s|\s\d{7,9}`),
			Colors: []string{tool.Red},
		},
		{
			// Size 'T'
			Regex:  regexp.MustCompile(`\s\d*[.,]?\dTi?\s|\s\d{10,12}`),
			Colors: []string{tool.Red},
		},
		{
			// Mounted on
			Regex:  regexp.MustCompile(`(?m)(/[-\w\d. ]*)+$`),
			Colors: []string{tool.Green},
		},
		{
			// 0-60%
			Regex:  regexp.MustCompile(`\s[1-6]?[0-9]%\s`),
			Colors: []string{tool.Green},
		},
		{
			// 70-89%
			Regex:  regexp.MustCompile(`\s[78][0-9]%\s`),
			Colors: []string{tool.Yellow},
		},
		{
			// 90-97%
			Regex:  regexp.MustCompile(`\s9[0-7]%\s`),
			Colors: []string{tool.Red},
		},
		{
			// 98-100%
			Regex:  regexp.MustCompile(`9[8-9]%|100%`),
			Colors: []string{tool.Red},
		},
		{
			// tmpfs lines
			Regex:  regexp.MustCompile(`(?m)^tmpfs.*`),
			Colors: []string{tool.Gray},
		},
	}
}
