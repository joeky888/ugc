package tool

import (
	"regexp"
)

type Conf struct {
	Regex      *regexp.Regexp
	RegexGroup string
	Colors     []string
}

const (
	Default = "%s"
	Black   = "\033[0;30m%s\033[0m"
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Blue    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"
	Gray    = "\033[1;90m%s\033[0m"
)
