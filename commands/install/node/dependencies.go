package node

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkDependencies() {
	output.Info("Checking all dependencies are installed by running 'yarn install'...")
	runner.CheckForMessage(
		"yarn install",
		"error",
		dependenciesInstalled,
		installDependencies,
	)
}

func dependenciesInstalled() error {
	output.Success("Node dependencies are installed.")
	return nil
}

func installDependencies() error {
	output.Fail("Node packages are missing.")
	output.Info("Attempting installation with:")
	output.Info("$ yarn install")
	return nil
}
