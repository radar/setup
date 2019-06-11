
package ruby

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/utils"
)

func checkBundler() error {
	if !utils.FileExists("Gemfile") {
		return nil
	}

	output.Info("Checking if Bundler is installed")
	runner.CheckForMessage(
		"gem list -i bundler",
		"false",
		bundlerInstalled,
		installBundler,
	)

	return nil
}

func bundlerInstalled() error {
	output.Success("Bundler is installed.")
	return nil
}

func installBundler() error {
	installCommand := "bundle install"

	output.Fail("Bundler is missing.")
	output.Info("Installing it with: ")
	output.Info("$ " + installCommand)
	runner.Stream(installCommand)
	return nil
}
