package elixir

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/utils"
)

func mixFileExists() bool {
	if utils.FileExists("mix.exs") {
		output.Success("mix.exs exists. Will attempt dependency installation.")
		return true
	} else {
		output.Skip("mix.exs does not exist. Skipping dependency installation.")
		return false
	}
}

func checkDependencies() {
	output.Info("Checking all dependencies are installed by running 'mix deps'...")
	runner.CheckForMessage(
		"mix deps",
		"the dependency is not available",
		dependenciesInstalled,
		installDependencies,
	)
}

func dependenciesInstalled() error {
	output.Success("Elixir dependencies are installed.")
	return nil
}

func installDependencies() error {
	output.Fail("Hex packages are missing.")
	output.Info("Attempting installation with:")
	runner.StreamWithInfo("mix hex.local --if-missing")
	runner.StreamWithInfo("mix deps.get")

	return nil
}

func mixCompile() {
	output.Info("Ensuring things are compiled...")
	output.Info("$ mix compile")
	runner.Stream("mix compile")
}
