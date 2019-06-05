package installer

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/toolversions"
	"github.com/radar/setup/commands/install/elixir"
	"github.com/radar/setup/commands/install/elm"
  "github.com/radar/setup/commands/install/node"
	"github.com/radar/setup/commands/install/ruby"
	"github.com/urfave/cli"
)

type installer func() error

func Run(c *cli.Context) error {
	versions, err := toolversions.Load()
	if err != nil {
		return err
	}

	output.Success("Found a .tool-versions file, will check those packages are installed...")

	installers := make(map[string]installer)
	installers["elixir"] = elixir.Run
	installers["elm"] = elm.Run
	installers["nodejs"] = node.Run
	installers["ruby"] = ruby.Run

	for k := range versions.Versions {
		if (installers[k] != nil) {
			err := installers[k]()
			if err != nil {
				return err
			}

		} else {
			output.Fail("I don't know how to install " + k + ". You're on your own!")
		}
	}

	output.Success("You're all good to go!")

	return nil
}
