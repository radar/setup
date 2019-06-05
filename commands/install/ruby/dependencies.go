package ruby

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkDependencies() {
	output.Info("Checking all Bundler dependencies are installed by running 'bundle check'...")
	runner.CheckForMessage(
		"bundle check",
		"The following gems are missing",
		dependenciesInstalled,
		installDependencies,
	)
}

func dependenciesInstalled() {
	output.Success("Bundler dependencies are installed.")
}

func installDependencies() {
	output.Fail("Gems are missing.")
	output.Info("Attempting installation with:")

	installCommand := "bundle install"
	output.Info("$ " + installCommand)
	runner.Stream(installCommand)
}
