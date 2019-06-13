package elixir

import (
	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
	"github.com/radar/setup/utils"
)

func mixFileExists() bool {
	if utils.FileExists("mix.exs") {
		output.Success("mix.exs exists. Will attempt dependency installation.", 4)
		return true
	} else {
		output.Skip("mix.exs does not exist. Skipping dependency installation.", 4)
		return false
	}
}

func checkDependencies() {
	output.Info("Checking all dependencies are installed by running 'mix deps'...", 6)
	runner.CheckForMessage(
		"mix deps",
		"the dependency is not available",
		dependenciesInstalled,
		installDependencies,
	)
}

func dependenciesInstalled() error {
	output.Success("Elixir dependencies are installed.", 6)
	return nil
}

func installDependencies() error {
	output.Fail("Hex packages are missing.", 6)
	output.Info("Attempting installation with:", 8)
	runner.StreamWithInfo("mix deps.get", 8)

	return nil
}

func mixCompile() {
	output.Info("Ensuring things are compiled...", 6)
	runner.StreamWithInfo("mix compile", 8)
}
