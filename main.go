package main

import (
	"github.com/radar/setup/installer"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "setup"
	app.Usage = "Setup your local development environment"
	app.Action = installer.Run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
