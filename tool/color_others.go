// +build !windows

package tool

const (
	Default = "%s"
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Blue    = "\033[1;36m%s\033[0m"
	Cyan    = Blue
	Gray    = "\033[1;90m%s\033[0m"
	// Black   = "\033[0;30m%s\033[0m"
	// White   = "\033[1;37m%s\033[0m"
)