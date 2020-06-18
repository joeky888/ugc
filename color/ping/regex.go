package ping

import (
	"regexp"

	"github.com/joeky888/ugc/tool"
)

var Config = []tool.Conf{
	{
		// IP
		Regex: regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`),
		Color: tool.Blue,
	},
	{
		// ipv6 number
		Regex: regexp.MustCompile(`(([0-9a-fA-F]{1,4})?\:\:?[0-9a-fA-F]{1,4})+`),
		Color: tool.Red,
	},
	{
		// icmp_seq=##
		Regex: regexp.MustCompile(`icmp_seq=(\d+)`),
		Color: tool.Yellow,
	},
	{
		// ttl=#
		Regex: regexp.MustCompile(`ttl=(\d+)`),
		Color: tool.Red,
	},
	{
		// name
		Regex: regexp.MustCompile(`(?:[fF]rom|PING)\s(\S+)\s`),
		Color: tool.Blue,
	},
	{
		// time
		Regex: regexp.MustCompile(`([0-9\.]+)\s?ms`),
		Color: tool.Green,
	},
	{
		// DUP
		Regex: regexp.MustCompile(`DUP\!`),
		Color: tool.Red,
	},
	{
		// OK
		Regex: regexp.MustCompile(`0(\.0)?% packet loss`),
		Color: tool.Green,
	},
	{
		// Errors
		Regex: regexp.MustCompile(`(Destination Host Unreachable|100(\.0)?% packet loss)`),
		Color: tool.Red,
	},
	{
		// unknown host
		Regex: regexp.MustCompile(`.+unknown\shost\s(.+)`),
		Color: tool.Red,
	},
	{
		// statistics header
		Regex: regexp.MustCompile(`--- (\S+) ping statistics ---`),
		Color: tool.Blue,
	},
	{
		// last line min/avg/max/mdev
		Regex: regexp.MustCompile(`\=\s([0-9\.]+)\/([0-9\.]+)\/([0-9\.]+)\/([0-9\.]+)`),
		Color: tool.Yellow,
	},
	{
		// these are good for nping
		Regex: regexp.MustCompile(`SENT|RCVD`),
		Color: tool.Red,
	},
	{
		// nping
		Regex: regexp.MustCompile(`unreachable`),
		Color: tool.Red,
	},
}
