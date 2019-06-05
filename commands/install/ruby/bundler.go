
package ruby

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkBundler() error {
	output.Info("Checking if Bundler is installed")
	runner.CheckForMessage(
		"gem list -i bundler",
		"false",
		bundlerInstalled,
		installBundler,
	)

	return nil
}

func bundlerInstalled() {
	output.Success("Bundler is installed.")
}

func installBundler() {
	installCommand := "bundle install"

	output.Fail("Bundler is missing.")
	output.Info("Installing it with: ")
	output.Info("$ " + installCommand)
	runner.Stream(installCommand)
}
