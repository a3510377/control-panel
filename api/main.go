package main

import (
	"os"

	"github.com/a3510377/control-panel/cli"
)

func main() {
	if err := cli.NewCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
