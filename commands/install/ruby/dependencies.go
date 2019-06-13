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


	output.Info("Checking all Bundler dependencies are installed by running 'bundle check'...", 4)
	runner.CheckForMessage(
		"bundle check",
		"The following gems are missing",
		dependenciesInstalled,
		installDependencies,
	)

	return nil
}

func dependenciesInstalled() error {
	output.Success("Bundler dependencies are installed.", 6)
	return nil
}

func installDependencies() error {
	output.Fail("Gems are missing.", 4)
	output.Info("Attempting installation with:", 6)

	installCommand := "bundle install"
	runner.StreamWithInfo(installCommand, 6)

	return nil
}
