package main

import (
	"fmt"
	"os"

	"github.com/radar/setup/installer"
	"github.com/urfave/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "setup"
	app.Usage = "Setup your local development environment"
	app.Action = installer.Run
	app.Version = fmt.Sprintf("%v, commit %v", version, commit)

	app.Run(os.Args)
}
