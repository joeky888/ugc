package ping

import (
	"regexp"

	"github.com/joeky888/ugc/tool"
)

// NewConfig ping regex config
func NewConfig() []tool.Conf {
	return []tool.Conf{
		{
			// IP
			Regex:  regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`),
			Colors: []string{tool.Blue},
		},
		{
			// ipv6 number
			Regex:  regexp.MustCompile(`(([0-9a-fA-F]{1,4})?\:\:?[0-9a-fA-F]{1,4})+`),
			Colors: []string{tool.Red},
		},
		{
			// icmp_seq=##
			Regex:  regexp.MustCompile(`icmp_seq=(\d+)`),
			Colors: []string{tool.Yellow},
		},
		{
			// ttl=#
			Regex:  regexp.MustCompile(`ttl=(\d+)`),
			Colors: []string{tool.Red},
		},
		{
			// name
			Regex:  regexp.MustCompile(`(?:[fF]rom|PING)\s(\S+)\s`),
			Colors: []string{tool.Blue},
		},
		{
			// last line min/avg/max/mdev
			Regex:  regexp.MustCompile(`(\s)(min|[0-9\.]+)(\/)(avg|[0-9\.]+)(\/)(max|[0-9\.]+)(\/)(mdev|[0-9\.]+)`),
			Colors: []string{tool.Default, tool.Yellow, tool.Default, tool.Blue, tool.Default, tool.Red, tool.Default, tool.Red},
		},
		{
			// time
			Regex:  regexp.MustCompile(`([0-9\.]+)\s?ms`),
			Colors: []string{tool.Green},
		},
		{
			// DUP
			Regex:  regexp.MustCompile(`DUP\!`),
			Colors: []string{tool.Red},
		},
		{
			// OK
			Regex:  regexp.MustCompile(`0(\.0)?% packet loss`),
			Colors: []string{tool.Green},
		},
		{
			// Errors
			Regex:  regexp.MustCompile(`(Destination Host Unreachable|100(\.0)?% packet loss)`),
			Colors: []string{tool.Red},
		},
		{
			// unknown host
			Regex:  regexp.MustCompile(`.+unknown\shost\s(.+)`),
			Colors: []string{tool.Red},
		},
		{
			// statistics header
			Regex:  regexp.MustCompile(`--- (\S+) ping statistics ---`),
			Colors: []string{tool.Blue},
		},
		{
			// these are good for nping
			Regex:  regexp.MustCompile(`SENT|RCVD`),
			Colors: []string{tool.Red},
		},
		{
			// nping
			Regex:  regexp.MustCompile(`unreachable`),
			Colors: []string{tool.Red},
		},
	}
}
