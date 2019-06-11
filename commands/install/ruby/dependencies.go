package ruby

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/utils"
)

func checkDependencies() error {
	if !utils.FileExists("Gemfile") {
		return nil
	}


	output.Info("Checking all Bundler dependencies are installed by running 'bundle check'...")
	runner.CheckForMessage(
		"bundle check",
		"The following gems are missing",
		dependenciesInstalled,
		installDependencies,
	)

	return nil
}

func dependenciesInstalled() error {
	output.Success("Bundler dependencies are installed.")
	return nil
}

func installDependencies() error {
	output.Fail("Gems are missing.")
	output.Info("Attempting installation with:")

	installCommand := "bundle install"
	output.Info("$ " + installCommand)
	runner.Stream(installCommand)

	return nil
}
