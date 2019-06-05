package main

import (
	"os"

	"github.com/radar/setup/installer"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "setup"
	app.Usage = "Setup your local development environment"
	app.Action = installer.Run

	app.Run(os.Args)
}
