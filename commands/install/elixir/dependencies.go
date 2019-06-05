package elixir

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkDependencies() {
	output.Info("Checking all dependencies are installed by running 'mix deps'...")
	runner.CheckForMessage(
		"mix deps",
		"the dependency is not available",
		dependenciesInstalled,
		installDependencies,
	)
}

func dependenciesInstalled() {
	output.Success("Elixir dependencies are installed.")
}

func installDependencies() {
	output.Fail("Hex packages are missing.")
	output.Info("Attempting installation with:")
	output.Info("$ mix hex.local --if-missing")
	output.Info("$ mix deps.get")

	runner.Stream("mix hex.local --if-missing")
	runner.Stream("mix deps.get")
}
