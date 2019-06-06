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
		Remedy: installViaASDF,
	}

	err := tool.Install()
	if err != nil {
		return err
	}

	checkDependencies()
	mixCompile()
	runEctoSetup()

	return nil
}

func installViaASDF() string {
	return "To fix this issue, you can run \"asdf install\"."
}
