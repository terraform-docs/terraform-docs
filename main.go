package main

import (
	"os"

	"github.com/segmentio/terraform-docs/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
