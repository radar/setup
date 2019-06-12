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

	fmt.Println(fmt.Sprintf("%v, commit %v, built at %v", version, commit, date))
	app := cli.NewApp()
	app.Name = "setup"
	app.Usage = "Setup your local development environment"
	app.Action = installer.Run

	app.Run(os.Args)
}
