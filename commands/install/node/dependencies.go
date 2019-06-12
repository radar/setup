package node

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkForYarn() (err error) {
	output.Info("Ensuring Yarn is installed")
	return runner.OptionalAction("which yarn", yarnInstalled, installYarn)
}

func yarnInstalled() error {
	output.Success("Yarn is installed.")
	return nil
}

func installYarn() error {
	output.Fail("Yarn is not installed.")
	output.Info("Installing yarn:")
	runner.StreamWithInfo("npm install -g yarn")
	return nil
}

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
	runner.StreamWithInfo("yarn install")
	return nil
}
