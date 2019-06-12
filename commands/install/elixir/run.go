package elixir

import (
	"github.com/radar/setup/tool"
)

const versionCommand = "elixir -v"

func Run() (err error) {
	tool := tool.Tool{
		Name: "Elixir",
		PackageName: "elixir",
		Executable: "elixir",
		VersionCommand: "elixir -v",
		VersionRegexp: `Elixir ([\d\.]{3,})`,
		Sources: []tool.Source{tool.ASDF},
	}

	err = tool.SetExpectedVersion()
	if err != nil {
		return err
	}

	err = checkForErlang()
	if err != nil {
		return err
	}

	err = tool.Install()
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
