package main

import (
	"log"

	"github.com/joeky888/ugc/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}
