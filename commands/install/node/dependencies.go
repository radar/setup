package node

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)


func checkForYarn() (err error) {
	output.Info("Ensuring Yarn is installed", 2)
	return runner.OptionalAction("which yarn", yarnInstalled, installYarn)
}

func yarnInstalled() error {
	output.Success("Yarn is installed.", 4)
	return nil
}

func installYarn() error {
	output.Fail("Yarn is not installed.", 4)
	output.Info("Installing yarn:", 6)
	runner.StreamWithInfo("npm install -g yarn", 6)
	return nil
}

func checkDependencies() {
	output.Info("Checking all dependencies are installed by running 'yarn install'...", 2)
	runner.CheckForMessage(
		"yarn install",
		"error",
		dependenciesInstalled,
		installDependencies,
	)
}

func dependenciesInstalled() error {
	output.Success("Node dependencies are installed.", 2)
	return nil
}

func installDependencies() error {
	output.Fail("Node packages are missing.", 4)
	output.Info("Attempting installation with:", 6)
	runner.StreamWithInfo("yarn install", 6)
	return nil
}
