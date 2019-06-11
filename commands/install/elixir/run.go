package elixir

import (
	"github.com/radar/setup/tool"
)

const versionCommand = "elixir -v"

func Run() error {
	tool := tool.Tool{
		Name: "Elixir",
		PackageName: "elixir",
		Executable: "elixir",
		VersionCommand: "elixir -v",
		VersionRegexp: `Elixir ([\d\.]{3,})`,
	}

	err := tool.Install()
	if err != nil {
		return err
	}

	ensureHexPresent()
	ensureRebarPresent()

	if mixFileExists() {
		checkDependencies()
		mixCompile()
		runEctoSetup()
	}

	return nil
}
