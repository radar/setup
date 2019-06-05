package installer

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/common/toolversions"
	"github.com/radar/setup/commands/install/elixir"
  "github.com/radar/setup/commands/install/node"
	"github.com/radar/setup/commands/install/ruby"
	"github.com/urfave/cli"
)

type installer func(c *cli.Context) error

func Run(c *cli.Context) error {
	versions, err := toolversions.Load()
	if err != nil {
		return err
	}

	output.Success("Found a .tool-versions file, will check those packages are installed...")

	installers := make(map[string]installer)
	installers["elixir"] = elixir.Run
	installers["nodejs"] = node.Run
	installers["ruby"] = ruby.Run



	for k := range versions.Versions {
		installers[k](c)
}

	return nil

}
